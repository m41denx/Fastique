import Barcode from "react-barcode";


export default function Ticket(props) {

    return <div className="ticket text-center rotate-6 border-solid border-2 border-[#9facbc]">
        <div className="p-6 bg-[#f4f5f6]">
            <div className="flex items-center justify-center gap-4 text-[#5e7186]">
                <img src={props.img} className="w-16" />
                <div>
                    <p className="text-4xl">{props.title}</p>
                    <p>{props.subtitle}</p>
                </div>
            </div>
        </div>
        <div className="p-4">
            <p className="text-6xl my-12">{props.ticker}</p>
            <div className="ticket__timing flex items-center justify-center">
                <p>
                    <span className="u-upper ticket__small-label">Дата</span>
                    <span className="ticket__detail">{props.date}</span>
                </p>
                <p>
                    <span className="u-upper ticket__small-label">Время</span>
                    <span className="ticket__detail">{props.time}</span>
                </p>
                <p>
                    <span className="u-upper ticket__small-label">Загруженность</span>
                    <span className="ticket__detail">{props.load}</span>
                </p>
            </div>
            <div className="flex items-center gap-2 flex-wrap mt-2">
                <p>Вид услуги</p>
                {props.labels?.map((v, i)=>
                    <span key={i} className="text-white bg-[#9facbc] rounded-full px-4 py-0.5">{v}</span>
                    )}
            </div>
            <p className="mt-4 text-[#5e7186]">{props.note}</p>
            <div className="flex justify-center">
                <Barcode value={props.ticker} displayValue={false} height={64} />
            </div>
        </div>
    </div>
}