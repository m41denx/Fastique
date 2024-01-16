import Ticket from "@/components/Ticket";
import {Accordion, AccordionItem} from "@szhsin/react-accordion";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {useEffect, useState} from "react";
import {Button, DatePicker, Form, Select, TimePicker} from "antd";
import useSWR from "swr";
import wapi from "@/config";
import axios from "axios";

const srv_labels = [
    {label: "В один конец", tag: "K"},
    {label: "Ну поплачь", tag: "P"}
]

const times = [
    {time: "10:00", taken: true},
    {time: "10:30", taken: false},
    {time: "11:00", taken: true},
    {time: "11:30", taken: true},
    {time: "12:00", taken: true},
    {time: "12:30", taken: false},
    {time: "13:00", taken: false},
    {time: "13:30", taken: true},
]

export function getServerSideProps(ctx) {

    return {
        props: {
            labels: srv_labels,
        }
    }
}


const months = ["Янв", "Фев", "Мар", "Апр", "Май", "Июн", "Июл", "Авг", "Сен", "Окт", "Ноя", "Дек"]




function daysInMonth (date, offset=0) { // Use 1 for January, 2 for February, etc.
    return new Date(date.getFullYear(), date.getMonth()+1+offset, 0).getDate();
}
export default function Order(props) {

    const [preform, setPreform] = useState({
        date: new Date(),
        time: times[0],
        branch: 0,
        label: 0
    })

    const {data: labels} = useSWR(wapi("fetch/labels"),
        (url) => axios.get(url)
            .then(r => r.data))

    const {data: branches} = useSWR(wapi(`fetch/branches/${preform.label}`),
        (url) => axios.get(url)
            .then(r => r.data))

    const [orgData, setOrgData] = useState({})
    useEffect(()=>{
        fetch("http://localhost:5000/fetch/org").then(r=>r.json()).then(j=>{
            setOrgData(j)
        })
    },[])



    return <div className="h-screen flex flex-col justify-center items-center">
        <div className="flex gap-24">
            {/*<Ticket img="https://img.icons8.com/pastel-glyph/128/secured-letter--v1.png"*/}
            {/*title={orgData.title} subtitle={orgData.desc} ticker={label.tag+"-07"}*/}
            {/*date={`${day} ${month}`} time={time.time} load={"Низкая"} labels={[label.label]}*/}
            {/*note="Адрес: ул Пушкина, 12, отделение №030"/>*/}

            <div className="flex flex-col w-96 h-full bg-[#f4f5f6] p-4 rounded-2xl border-solid border-2 border-[#9facbc]">
                <p className="bg-slate-600 px-4 py-2 rounded-xl text-white text-3xl">Тут заголовок</p>

               <Form name="basic"
                      labelCol={{span: 6}}
                      wrapperCol={{span: 24}}
                      initialValues={{remember: true}}
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
                            label: v
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
                           branch: v
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
                        <DatePicker className="w-full" placeholder="Укажите дату" format="DD.MM.YYYY" disabledDate={d=>d<new Date().setDate(new Date().getDate()-1)} />
                    </Form.Item>
                    <Form.Item
                        label="Время"
                        name="time"
                        rules={[{
                            required: true,
                            message: 'Укажите время',
                        }]}>
                        <TimePicker format="HH:mm" placeholder="Укажите время" minuteStep={30} className="w-full" />
                    </Form.Item>

                    <Form.Item className="flex justify-end">
                        <Button type="primary" htmlType="submit" className="bg-blue-600">
                            Забронировать
                        </Button>
                    </Form.Item>
               </Form>
            </div>
        </div>
    </div>
}