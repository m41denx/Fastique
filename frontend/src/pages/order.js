import Ticket from "@/components/Ticket";
import {Accordion, AccordionItem} from "@szhsin/react-accordion";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {useState} from "react";

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
export default function Order({labels}) {

    const date =  new Date()

    const offset = new Date(date.getFullYear(), date.getMonth(), 0).getDay()

    const weeknd = [""]
    const blook = (()=>{

        let b = []
        for(let i=-offset+1; i<=daysInMonth(date); i++)
            b.push(i<=0 ? daysInMonth(date,-1)+i : i)
        return b
    })()

    const [month, setMonth] = useState(months[date.getMonth()])
    const [day, setDay] = useState(date.getDate())
    const [label, setLabel] = useState(labels[0])
    const [time, setTime] = useState(times[0])

    return <div className="h-screen flex flex-col justify-center items-center">
        <div className="flex gap-24">
            <Ticket img="https://img.icons8.com/pastel-glyph/128/secured-letter--v1.png"
            title="Билет в Бобруйск" subtitle={"Это надолго"} ticker={label.tag+"-07"}
            date={`${day} ${month}`} time={time.time} load={"Низкая"} labels={[label.label]}
            note="Адрес: ул Пушкина, 12, отделение №030"/>


            <div className="flex flex-col w-96 h-full bg-[#f4f5f6] p-4 rounded-2xl  border-solid border-2 border-[#9facbc]">
                <p className="bg-slate-600 px-4 py-2 rounded-xl text-white text-3xl">Тут заголовок</p>

                <Accordion>

                    <AccordionItem initialEntered className="rounded-xl mt-4 bg-[#9facbc] bg-opacity-75 p-2" header={
                        <p className="flex items-center justify-between">
                                <span className="text-lg rounded-lg bg-[#f4f5f6] px-2">Выберите услугу</span>
                                <FontAwesomeIcon icon={faChevronDown} className="group px-2" />
                            </p>
                    } buttonProps={{
                        className: ({isEnter})=> `w-full ${isEnter? "group:rotate-180": ""}`
                    }}>
                        <div className="flex flex-wrap gap-2 mt-2">
                            {labels.map((lbl, i) => (
                                <p key={i} className="px-2 text-lg rounded-lg bg-[#f4f5f6] text-center bg-opacity-100 cursor-pointer hover:!bg-gray-300 hover:!text-black"
                                   style={lbl===label?{
                                       backgroundColor: "green",
                                       color: "white",
                                   }:{}} onClick={()=>setLabel(lbl)}>{lbl.label}</p>
                            ))}
                        </div>
                    </AccordionItem>

                    <AccordionItem className="rounded-xl mt-4 bg-[#9facbc] bg-opacity-75 p-2" header={
                        <>
                            <p className="flex items-center justify-between">
                                <span className="text-lg rounded-lg bg-[#f4f5f6] px-2">Выберите дату</span>
                                <FontAwesomeIcon icon={faChevronDown} className="group px-2" />
                            </p>
                        </>
                    } buttonProps={{
                        className: ({isEnter})=> `w-full ${isEnter? "group:rotate-180": ""}`
                    }}>
                        <div className="grid grid-cols-7 gap-1 mt-2">
                            {blook.map((tday, i) => (
                                <p key={i} className="text-lg rounded-lg bg-[#f4f5f6] text-center bg-opacity-100 cursor-pointer hover:!bg-gray-300 hover:!text-black"
                                   style={
                                       (i===day+offset-1) ? {
                                           backgroundColor: "#0000cc",
                                           color: "white",
                                       }
                                       : (
                                           (i%7===6 || i%7===5) ? {
                                               backgroundColor: "rgb(220 38 38)",
                                               color: "white",
                                           }
                                           : {}
                                           )
                                } onClick={()=>setDay(i-offset+1)}>{tday}</p>
                            ))}
                        </div>
                        <div className="grid grid-cols-5 gap-1 mt-4">
                            {times.map((tme, i) => (
                                <p key={i} className="text-lg rounded-lg bg-[#f4f5f6] text-center bg-opacity-100 cursor-pointer hover:!bg-gray-300 hover:!text-black"
                                   style={
                                       {
                                           backgroundColor: tme.taken?"rgb(220 38 38)":"green",
                                           color: "white",
                                       }
                                   } onClick={()=>setTime(tme)}>{tme.time}</p>
                            ))}
                        </div>
                    </AccordionItem>
                </Accordion>

                <p className="text-white bg-blue-600 px-4 py-2 rounded-xl text-center text-lg mt-auto cursor-pointer hover:bg-blue-800">
                    Та самая кнопка
                </p>
            </div>
        </div>
    </div>
}