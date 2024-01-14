import {Breadcrumb, List} from "antd";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faHome} from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";

export default function Screen(props) {

    const locs = [{label:"loc1"}, {label:"loc2"}, {label:"loc3"}]
    const wins = [{label:"win1"}, {label:"win2"}, {label:"win3"}]

    const queue = [{name:"K-11", time: "10:10"},{name:"K-11", time: "10:10"}]

    return <div className="h-screen">
        <div className="flex items-center text-xl bg-gray-200 shadow drop-shadow p-2">
            <Breadcrumb items={[
                {title: <Link href="/"><FontAwesomeIcon icon={faHome}/></Link>},
                {title: "loc1", menu: {items: locs}},
                {title: "win1", menu: {items: locs}}
            ]} />
        </div>
        <div className="grid grid-cols-3 p-4 gap-4 h-5/6">
            <div className="col-span-2 flex flex-col items-center rounded-xl bg-gray-200 h-full p-4 shadow drop-shadow">
                <p className="text-4xl">Окно 9</p>
                <p className="text-[8rem] my-auto">Л-11</p>
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