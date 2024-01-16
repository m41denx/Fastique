import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import axios from "axios";
import wapi from "@/config";
import toast, {Toaster} from "react-hot-toast";
import {Breadcrumb, Button, Dropdown, List} from "antd";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown, faHome} from "@fortawesome/free-solid-svg-icons";
import NavBar from "@/components/NavBar";
import useSWR, {mutate} from "swr";
import Link from "next/link";


const getFromListMap = (list, id, field='ID') => {
    for (let i = 0; i < list.length; i++) {
        if (list[i][field] == id) {
            return list[i]
        }
    }
    return {}
}

const getOrderFromList = (list, id, field='ID') => {
    for (let i = 0; i < list.length; i++) {
        if (list[i][field] == id) {
            return i
        }
    }
    return -1
}

export default function Process(props) {
    const router = useRouter()
    const [userData, setUserData] = useState({})


    const {data: branches} = useSWR(wapi("fetch/branches"),
        (url) => axios.get(url)
            .then(r => r.data))



    const [selected, setSelected] = useState({
        branch: branches?branches[0]:{name:"Не выбрано"},
        sp: branches?branches[0]?.service_points[0]||{name: "Не выбрано"}:{name:"Не выбрано"},
        authed: false
    })

    const [current, setCurrent] = useState(null)


    const {data: tickets} = useSWR(wapi(`worker/queue/${selected.sp?.ID}`),
        (url) => axios.get(url, {headers: {Authorization: props.token}})
            .then(r => r.data))

    const queue = tickets?.map((m, id)=> {
        const d = new Date(m.beginTime)
        if (new Date(m.endTime).getFullYear() > 1970)
            return {}
        return {
            name: m.name,
            time: `${("0"+d.getHours()).slice(-2)}:${("0"+d.getMinutes()).slice(-2)}`,
            t: m
        }
    })||[]

    const claimSP = () => {
        axios.post(wapi(`worker/b/${selected.sp.ID}`), {}, {headers: {Authorization: props.token}}).then(r => {
            setSelected({...selected, authed: true})
        }).catch(e=>{
            toast.error("Не удалось занять точку обслуживания")
        })
    }
    const unclaimSP = () => {
        axios.delete(wapi(`worker/b/${selected.sp.ID}`), {headers: {Authorization: props.token}}).then(r => {
            setSelected({...selected, authed: false})
        }).catch(e=>{
            toast.error("Не удалось выйти из точки обслуживания")
        })
    }

    const acceptTicket = (t) => {
        axios.post(wapi(`worker/claim/${selected.sp.ID}/${t.ID}`), {}, {headers: {Authorization: props.token}}).then(r => {
            toast.success("Тикет принят")
            mutate(wapi(`worker/queue/${selected.sp.ID}`))
            setCurrent(t)
        }).catch(e=>{
            toast.error("Не удалось взять тикет")
        })
    }

    const declineTicket = (t) => {
        axios.delete(wapi(`worker/claim/${selected.sp.ID}/${t.ID}`), {headers: {Authorization: props.token}}).then(r => {
            toast.success("Тикет помечен как завершенный")
            mutate(wapi(`worker/queue/${selected.sp.ID}`))
            setCurrent(null)
        }).catch(e=>{
            toast.error("Не удалось пометить тикет как завершенный")
        })
    }

    return <div className="h-screen">
        <Toaster />
        <NavBar setUserData={setUserData} userData={userData} token={props.token} />
        <div className="p-4 h-full grid grid-cols-3 gap-4">
            <div className="col-span-2 flex flex-col gap-4">
                <div className="bg-gray-200 rounded-xl p-4 flex items-center gap-4">
                    <Breadcrumb items={[
                        {title: <Link href="/"><FontAwesomeIcon icon={faHome}/></Link>},
                        {title: selected.branch?.name||'-', menu: {items: branches?.map((b)=>{
                                    return {title: b.name, onClick: ()=>{setSelected({...selected, branch: b})}}
                                })}},
                        {title: selected.branch?.service_points?`Окно ${getOrderFromList(selected.branch.service_points, selected.sp.ID) + 1}`:'-',
                            menu: {items: selected.branch?.service_points?.map((b,i)=>{
                                    return {title: `Окно ${i+1}`, onClick: ()=>{setSelected({...selected, sp: b})}}
                                })||[]}}
                    ]} />
                    <Button type="primary" danger={selected.authed} onClick={()=>{
                        if(!selected.authed)
                            claimSP()
                        else
                            unclaimSP()
                    }}>{selected.authed?"Выйти":"Активировать"}</Button>
                </div>
                {current&&<div className="bg-gray-200 rounded-xl p-4 flex justify-center flex-col items-center gap-4">
                    <p className="text-2xl">{current.name}</p>
                    <p className="flex gap-2 items-center">
                        <span className="bg-white rounded-xl p-2">{("0"+new Date(current.beginTime).getHours()).slice(-2)}:{("0"+new Date(current.beginTime).getMinutes()).slice(-2)}</span>
                        <span className="bg-white rounded-full h-2 w-2 block"></span>
                        <span className="bg-white rounded-full h-2 w-2 block"></span>
                        <span className="bg-white rounded-full h-2 w-2 block"></span>
                        <span className="bg-white rounded-xl p-2">{("0"+new Date().getHours()).slice(-2)}:{("0"+new Date().getMinutes()).slice(-2)}</span>
                    </p>
                    <div className="flex gap-4 mt-4">
                        <Button className="flex-1" size="large" type="primary" onClick={()=>declineTicket(current)}>Завершить</Button>
                    </div>
                </div>}
            </div>
            <div className="bg-gray-200 rounded-xl p-4 overflow-y-scroll">
                <p className="text-xl">Очередь</p>
                <List dataSource={queue} renderItem={(item, i) => item.name&&<List.Item>
                    <div className="bg-white rounded-xl p-2 w-full">
                       <div className="flex items-center justify-between gap-4">
                           <span className="text-2xl w-12 bg-gray-200 h-12 rounded-full flex items-center justify-center">{i+1}</span>
                           <span className="text-4xl">{item.name}</span>
                           <span className="text-lg bg-gray-200 rounded-md px-2">{item.time}</span>
                       </div>
                        <div className="flex gap-4 mt-4">
                            <Button className="flex-1" type="primary" onClick={()=>acceptTicket(item.t)}>Принять</Button>
                            <Button className="flex-1" type="primary" danger onClick={()=>declineTicket(item.t)}>Отменить</Button>
                        </div>
                    </div>
                </List.Item>} />
            </div>
        </div>
    </div>
}