

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
            <img className="rounded-lg mx-auto" src="data:image/png;base64,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" alt="Fake barcode" />
        </div>
    </div>
}