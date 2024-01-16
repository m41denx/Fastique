import {useEffect, useState} from "react";
import useSWR from "swr";
import wapi from "@/config";
import axios from "axios";
import {Breadcrumb, Menu} from "antd";
import Link from "next/link";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHome} from "@fortawesome/free-solid-svg-icons";
import Ticket from "@/components/Ticket";


const getLoad = (n) => {
    switch (n){
        case -1:
            return "-"
        case 0:
        case 1:
            return "Низкая"
        case 2:
        case 3:
            return "Средняя"
        default:
            return "Высокая"
    }
}

const months = ["Янв", "Фев", "Мар", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"]


const getFromListMap = (list, id, field='ID') => {
    for (let i = 0; i < list.length; i++) {
        if (list[i][field] == id) {
            return list[i]
        }
    }
    return {}
}

const renderTicker = (ticker, id) => {
    ticker = ticker||"#"
    return ticker.replace("#", id)
}


export default function OrderLocal(props) {

    const [orgData, setOrgData] = useState({})
    useEffect(()=>{
        fetch(wapi("fetch/org")).then(r=>r.json()).then(j=>{
            setOrgData(j)
        })
    },[])

    const {data: branches} = useSWR(wapi("fetch/branches"),
        (url) => axios.get(url)
            .then(r => r.data))

    const {data: labels} = useSWR(wapi("fetch/labels"),
        (url) => axios.get(url)
            .then(r => r.data))

    const [preform, setPreform] = useState({
        date: new Date(),
        time: new Date(),
        branch: branches&&branches[0],
        label: {}
    })

    const [result, setResult] = useState({
        order: "#",
        load: -1
    })

    const reserve = (label) => {
        let date = new Date()
        let req = {
            label: label.ID,
            branch: preform.branch.ID,
            date: preform.date.toISOString(),
        }
        axios.post(wapi("reserve"), req)
            .then(r => {
                setResult(r.data)
                setPreform({...preform, date: date, time: date, label: label})
            })
    }

    return <>
        <div className="flex items-center text-xl bg-gray-200 shadow drop-shadow p-2">
            <Breadcrumb items={[
                {title: <Link href="/"><FontAwesomeIcon icon={faHome}/></Link>},
                {title: preform.branch?.name||"Не выбрано", menu: {items: branches?.map((b)=>{
                            return {title: b.name, onClick: ()=>{setPreform({...preform, branch: b})}}
                        })}},
            ]} />
        </div>
        <div className="h-screen flex flex-col justify-center items-center">
            {preform.branch&&<div className="flex gap-24">

                <Ticket img="https://img.icons8.com/pastel-glyph/128/secured-letter--v1.png"
                        title={orgData.title} subtitle={orgData.desc} ticker={renderTicker(preform.label.template, result.order)}
                        date={`${preform.date.getDate()} ${months[preform.date.getMonth()]}`}
                        time={`${("0"+preform.time.getHours()).slice(-2)}:${("0"+preform.time.getMinutes()).slice(-2)}`}
                        load={getLoad(result.load)} labels={[preform.label.name]}
                        note={preform.branch.name}/>

                <div className="bg-gray-200 p-4 rounded-xl h-fit">
                    <p className="text-xl">Выберите услугу</p>
                    <Menu className="rounded-lg mt-2" items={labels?.map((l)=>{
                        return {key: l.ID, label: l.name, onClick: ()=>{
                            reserve(l)
                        }}
                    })}/>
                </div>
            </div>}
        </div>
        </>
}