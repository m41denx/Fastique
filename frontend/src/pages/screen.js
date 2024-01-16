import {Breadcrumb, List} from "antd";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHome} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";
import useSWR from "swr";
import axios from "axios";
import wapi, {WS_URL} from "@/config";
import {useState} from "react";
import useWebSocket from "react-use-websocket";


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

export default function Screen(props) {

    const {data: branches} = useSWR(wapi("fetch/branches"),
        (url) => axios.get(url)
            .then(r => r.data))

    const [selected, setSelected] = useState({
        branch: branches?branches[0]:{name:"Не выбрано"},
    })

    const {sendJsonMessage: sendMsg, lastJsonMessage: lastMsg, readyState} = useWebSocket(`${WS_URL}/${selected.branch?.ID}`)


    const queue = lastMsg?.map((m, id)=> {
        const d = new Date(m.beginTime)
        return {
            name: m.name,
            time: `${("0"+d.getHours()).slice(-2)}:${("0"+d.getMinutes()).slice(-2)}`,
        }
    })||[]

    return <div className="h-screen">
        <div className="flex items-center text-xl bg-gray-200 shadow drop-shadow p-2">
            <Breadcrumb items={[
                {title: <Link href="/"><FontAwesomeIcon icon={faHome}/></Link>},
                {title: selected.branch.name, menu: {items: branches?.map((b)=>{
                            return {title: b.name, onClick: ()=>{setSelected({...selected, branch: b})}}
                        })}},
            ]} />
        </div>
        <div className="grid grid-cols-3 p-4 gap-4 h-5/6">
            <div className="col-span-2 flex flex-col items-center rounded-xl bg-gray-200 h-full p-4 shadow drop-shadow">
                <p className="text-4xl">{(lastMsg&&lastMsg.length&&lastMsg[0].spid)?
                    `Окно ${getOrderFromList(selected.branch.service_points, lastMsg[0].spid)+1}`
                    :"-"}</p>
                <p className="text-[8rem] my-auto">{(lastMsg&&lastMsg.length&&lastMsg[0].spid)?lastMsg[0].name:"⌛"}</p>
            </div>
            <div className="overflow-y-hidden rounded-xl bg-gray-200 h-full p-4 shadow drop-shadow">
                <p className="text-xl">Очередь</p>
                <List dataSource={queue} renderItem={(item, i) => <List.Item>
                    <div className="bg-white rounded-xl p-2 flex items-center justify-between w-full gap-4">
                        <span className="text-2xl w-12 bg-gray-200 h-12 rounded-full flex items-center justify-center">{i+1}</span>
                        <span className="text-4xl">{item.name}</span>
                        <span className="text-lg bg-gray-200 rounded-md px-2">{item.time}</span>
                    </div>
                </List.Item>} />
            </div>
        </div>
    </div>
}