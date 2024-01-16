import {useState} from "react";
import useSWR from "swr";
import wapi from "@/config";
import axios from "axios";
import {Breadcrumb} from "antd";
import Link from "next/link";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHome} from "@fortawesome/free-solid-svg-icons";


export default function OrderLocal(props) {

    const {data: labels} = useSWR(wapi("fetch/labels"),
        (url) => axios.get(url)
            .then(r => r.data))

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

    const locs = []

    return <>
        <div className="flex items-center text-xl bg-gray-200 shadow drop-shadow p-2">
            <Breadcrumb items={[
                {title: <Link href="/"><FontAwesomeIcon icon={faHome}/></Link>},
                {title: "loc1", menu: {items: locs}},
                {title: "win1", menu: {items: locs}}
            ]} />
        </div>
        <div className="h-screen flex flex-col justify-center items-center">

        <div className="flex gap-24"></div>
        </div>
        </>
}