import Ticket from "@/components/Ticket";
import {Accordion, AccordionItem} from "@szhsin/react-accordion";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {useEffect, useState} from "react";
import {Button, DatePicker, Form, Select, TimePicker} from "antd";
import useSWR from "swr";
import wapi from "@/config";
import axios from "axios";
import dayjs from "dayjs";

const range = (start, end) => {
    const result = [];
    for (let i = start; i < end; i++) {
        result.push(i);
    }
    return result;
};

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

function daysInMonth (date, offset=0) { // Use 1 for January, 2 for February, etc.
    return new Date(date.getFullYear(), date.getMonth()+1+offset, 0).getDate();
}
export default function Order(props) {

    const [preform, setPreform] = useState({
        date: new Date(),
        time: new Date(),
        branch: {},
        label: {}
    })

    const [result, setResult] = useState({
        order: "#",
        load: -1
    })

    const {data: labels} = useSWR(wapi("fetch/labels"),
        (url) => axios.get(url)
            .then(r => r.data))

    const {data: branches} = useSWR(wapi(`fetch/branches/${preform.label.ID}`),
        (url) => axios.get(url)
            .then(r => r.data))

    const [orgData, setOrgData] = useState({})
    useEffect(()=>{
        fetch(wapi("fetch/org")).then(r=>r.json()).then(j=>{
            setOrgData(j)
        })
    },[])

    const reserve = (d) => {
        let date = preform.date
        date.setHours(preform.time.getHours())
        date.setMinutes(preform.time.getMinutes())
        let req = {
            label: preform.label.ID,
            branch: preform.branch.ID,
            date: preform.date.toISOString(),
        }
        axios.post(wapi("reserve"), req)
            .then(r => {
                setResult(r.data)
            })
    }


    return <div className="h-screen flex flex-col justify-center items-center">
        <div className="flex gap-24">
            <Ticket img="https://img.icons8.com/pastel-glyph/128/secured-letter--v1.png"
            title={orgData.title} subtitle={orgData.desc} ticker={renderTicker(preform.label.template, result.order)}
            date={`${preform.date.getDate()} ${months[preform.date.getMonth()]}`}
                    time={`${("0"+preform.time.getHours()).slice(-2)}:${("0"+preform.time.getMinutes()).slice(-2)}`}
                    load={getLoad(result.load)} labels={[preform.label.name]}
            note={preform.branch.name}/>

            {result.load<0 ? <div className="flex flex-col w-96 h-full bg-[#f4f5f6] p-4 rounded-2xl border-solid border-2 border-[#9facbc]">
                <p className="bg-slate-600 px-4 py-2 rounded-xl text-white text-2xl">{orgData.title}</p>

               <Form name="basic"
                      labelCol={{span: 6}}
                      wrapperCol={{span: 24}}
                      initialValues={{remember: true}}
                     onFinish={reserve}
                      autoComplete="off" className="p-4 bg-gray-200 rounded-xl mt-2">
                    <Form.Item
                        label="Услуга"
                        name="label"
                        rules={[{
                            required: true,
                            message: 'Выберите услугу',
                        }]}>
                        <Select placeholder="Выберите услугу" onChange={(v)=>setPreform({
                            ...preform,
                            label: getFromListMap(labels, v)
                        })}
                                options={labels?.map(lbl=>({label: lbl.name, value: lbl.ID}))} />
                    </Form.Item>
                   <Form.Item
                       label="Адрес"
                       name="branch"
                       rules={[{
                           required: true,
                           message: 'Выберите адрес',
                       }]}>
                       <Select placeholder="Выберите адрес"  onChange={(v)=>setPreform({
                           ...preform,
                           branch: getFromListMap(branches, v)
                       })}
                               options={branches?.map(lbl=>({label: lbl.name, value: lbl.ID}))} />
                   </Form.Item>
                    <Form.Item
                        label="Дата"
                        name="date"
                        rules={[{
                            required: true,
                            message: 'Укажите дату',
                        }]}>
                        <DatePicker className="w-full" placeholder="Укажите дату" format="DD.MM.YYYY" onChange={(v)=>setPreform({
                            ...preform,
                            date: dayjs(v).toDate()
                        })}
                                    disabledDate={d=>d<new Date().setDate(new Date().getDate()-1)}/>
                    </Form.Item>
                    <Form.Item
                        label="Время"
                        name="time"
                        rules={[{
                            required: true,
                            message: 'Укажите время',
                        }]}>
                        <TimePicker format="HH:mm" placeholder="Укажите время" minuteStep={30} hideDisabledOptions
                                    className="w-full" disabledTime={()=>({
                             disabledHours: ()=> new Date().getDate() == preform.date.getDate() ? range(0, new Date().getHours()) : [],
                             disabledMinutes: (v)=> v==new Date().getHours() ? range(0, new Date().getMinutes()) : []
                        })} onChange={(v)=>setPreform({...preform, time: dayjs(v).toDate()})} />
                    </Form.Item>

                    <Form.Item className="flex justify-end">
                        <Button type="primary" htmlType="submit" className="bg-blue-600">
                            Забронировать
                        </Button>
                    </Form.Item>
               </Form>
            </div>
                :<div className="bg-gray-200 rounded-xl p-4 h-fit">🎉 Вы успешено забронировали время</div> }
        </div>
    </div>
}