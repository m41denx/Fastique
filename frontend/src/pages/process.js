import {useEffect, useState} from "react";
import {useRouter} from "next/router";
import axios from "axios";
import wapi from "@/config";
import toast, {Toaster} from "react-hot-toast";
import {Dropdown} from "antd";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown} from "@fortawesome/free-solid-svg-icons";
import NavBar from "@/components/NavBar";


export default function Process(props) {

    const router = useRouter()
    const [userData, setUserData] = useState({})

    return <div className="h-screen">
        <Toaster />
        <NavBar setUserData={setUserData} userData={userData} token={props.token} />
        <div className="p-4 h-full grid grid-cols-3 gap-4">
            <div className="col-span-2">
                <div className="bg-gray-200 rounded-xl p-4"></div>
            </div>
            <div className="bg-gray-200 rounded-xl p-4"></div>
        </div>
    </div>
}