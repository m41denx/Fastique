import {Dropdown} from "antd";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {useEffect, useState} from "react";
import axios from "axios";
import wapi from "@/config";
import toast from "react-hot-toast";
import {useRouter} from "next/router";


export default function NavBar({userData, setUserData, token}) {
    const router = useRouter()
    const [orgData, setOrgData] = useState({})
    const privs = userData.role?.privileges||{}

    useEffect(() => {
        axios.get(wapi("fetch/org"), {validateStatus: () => true}).then(j => {
            setOrgData(j.data)
        }).catch(err => {
            toast.error("Ошибка соединения")
        })
    }, [])

    useEffect(() => {
        if (token===".") {
            router.push("/login")
            return
        }
        axios.get(wapi("auth/user"), {headers: {Authorization: token}, validateStatus: () => true}).then(j => {
            if (j.status === 200) {
                setUserData(j.data)
            } else {
                token&&router.push("/login")
            }
        }).catch(err => {
            toast.error("Ошибка соединения")
        })
    }, [token])



    return <div className="flex items-center justify-between text-lg bg-gray-200 shadow drop-shadow p-2">
        <span>{orgData.title}</span>
        <Dropdown menu={{items: [
                {key: "process", label: "Рабочее место", onClick: ()=>{router.push("/process")}},
                privs['admin'] && {key: "admin", label: "Администрирование", onClick: ()=>{router.push("/admin")}},
                {key: "logout", label: "Выход", onClick: ()=>{localStorage.removeItem("token"); router.push("/login")}},
            ]}} placement="bottomRight" arrow>
            <span className="cursor-pointer rounded-lg hover:bg-white px-2 py-1">{userData.username} <FontAwesomeIcon className="h-4" icon={faChevronDown}/></span>
        </Dropdown>
    </div>
}