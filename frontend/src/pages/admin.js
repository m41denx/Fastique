import {useRouter} from "next/router";
import NavBar from "@/components/NavBar";
import {useState} from "react";
import {Button, Form, Input, Menu, Modal, Select, Table} from "antd";
import {faAdd, faKey, faList, faPencil, faShop, faTicket, faTrash, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import useSWR, {mutate} from "swr";
import wapi from "@/config";
import axios from "axios";
import toast, {Toaster} from "react-hot-toast";

export default function Admin(props) {

    const router = useRouter()
    const [userData, setUserData] = useState({})
    const [page, setPage] = useState("users")

    const Page = (name, id, icon=faList) => {
        return <p className={`p-2 rounded-lg hover:bg-white transition-all duration-150 cursor-pointer ${page === id ? "bg-white" : ""}`}
                  onClick={() => setPage(id)}>
            <FontAwesomeIcon icon={icon} className="mr-2 w-6" />{name}
        </p>
    }

    return (
        <div className="h-screen">
            <Toaster />
            <NavBar setUserData={setUserData} userData={userData} token={props.token} />
            <div className="h-full grid grid-cols-4 gap-4 p-4">
                <div className="bg-gray-200 rounded-xl p-2 h-fit flex flex-col gap-2">
                    {Page("Сотрудники", "users", faUser)}
                    {Page("Роли", "roles", faKey)}
                    {Page("Точки обслуживания", "sp", faShop)}
                    {Page("Все заявки", "tickets", faTicket)}
                </div>
                <div className="col-span-3 bg-gray-200 rounded-xl p-2 h-fit">
                    {page === "users" && <UsersView token={props.token} />}
                    {page === "roles" && <RolesView token={props.token} />}
                    {page === "sp" && <BranchesView token={props.token} />}
                    {page === "tickets" && <div>Tickets</div>}
                </div>
            </div>
        </div>
    )
}

const RolesView = (props) => {
    const {data, error, isLoading} = useSWR(wapi("adm/roles"),
        (url) => axios.get(url, {headers: {"Authorization": props.token}})
            .then(r => r.data))

    const columns = [
        {title: "ID", dataIndex: "ID", key: "id"},
        {title: "Название", dataIndex: "name", key: "name"},
        {title: "Метаданные", dataIndex: "meta", key: "meta", render: (lbls) =>
                <p className="flex gap-1">{Object.keys(lbls||{}).map((k) =>
                    <span className="bg-gray-200 px-1 py-0.5 rounded-md" key={k}>{k}: <span className="bg-white rounded-sm px-1">{lbls[k]}</span></span>
                )}</p>
        },
        {title: "Права", dataIndex: "privileges", key: "privileges", render: (privs) =>
                <p className="flex gap-1">{Object.keys(privs).map((k) =>
                privs[k]&&<span className="bg-gray-200 px-1 py-0.5 rounded-md" key={k}>{k}</span>
                    )}</p>
        },
    ]
    return <div className="flex flex-col gap-2">
        <p className="text-xl">Роли</p>
        <Table columns={columns} dataSource={data||[]} />
        <Button type="primary" className="bg-blue-600 w-fit ml-auto">Создать роль</Button>
    </div>
}

const UsersView = (props) => {
    const {data, error, isLoading} = useSWR(wapi("adm/users"),
        (url) => axios.get(url, {headers: {"Authorization": props.token}})
            .then(r => r.data))

    const columns = [
        {title: "ID", dataIndex: "ID", key: "id"},
        {title: "Имя", dataIndex: "username", key: "username"},
        {title: "Метаданные", dataIndex: "meta", key: "meta", render: (lbls) =>
                <p className="flex gap-1">{Object.keys(lbls||{}).map((k) =>
                    <span className="bg-gray-200 px-1 py-0.5 rounded-md" key={k}>{k}: <span className="bg-white rounded-sm px-1">{lbls[k]}</span></span>
                )}</p>
        },
        {title: "Роль", dataIndex: "role", key: "role", render: (role) =>
                role&&<span className="bg-gray-200 px-1 py-0.5 rounded-md">{role.name}</span>
        },
        {key: "action", render: (_, user)=><div className="flex gap-1 w-fit">
                <Button type="primary" className="bg-blue-600 w-fit ml-auto"><FontAwesomeIcon icon={faPencil} /></Button>
                <Button type="primary" danger className="bg-red-600 w-fit ml-auto"><FontAwesomeIcon icon={faTrash} /></Button>
            </div>,
        width: "8rem"}
    ]

    return <div className="flex flex-col gap-2">
        <p className="text-xl">Сотрудники</p>
        <Table columns={columns} dataSource={data||[]} />
        <Button type="primary" className="bg-blue-600 w-fit ml-auto">Добавить сотрудника</Button>
    </div>
}


const BranchesView = (props) => {

    const [mopen, setMopen] = useState(false)
    const [selectedBranch, setSelectedBranch] = useState(null)

    const {data, error, isLoading} = useSWR(wapi("adm/branches"),
        (url) => axios.get(url, {headers: {"Authorization": props.token}})
            .then(r => r.data))

    const {data: labels} = useSWR(wapi("adm/labels"),
        (url) => axios.get(url, {headers: {"Authorization": props.token}})
            .then(r => r.data))


    const branchView = (branch)=> {
        let sps = branch.service_points.map((sp, i) => {
            return {
                key: sp.ID,
                icon: <FontAwesomeIcon icon={faTicket} className={sp.available?"!text-green-600":"!text-red-600"} />,
                label: <span className="flex items-center justify-between">
                    <span>Точка {i+1}</span>
                    <p className="flex gap-1">{sp.labels.map(lbl=>
                        <span className="bg-gray-200 px-1 py-0.5 rounded-md text-sm" key={lbl.ID}>{lbl.name}</span>
                    )}</p>
                    {sp.user.ID?
                        <span>
                        <FontAwesomeIcon icon={faUser} /> {sp.user?.username}
                        </span>:<span></span>
                    }

                    <Button type="primary" className="bg-blue-600 w-fit h-8">
                        <FontAwesomeIcon icon={faPencil}  className="h-3"/>
                    </Button>
                </span>
            }
        })
        sps.push({key: "add", icon: <FontAwesomeIcon icon={faAdd} />, label: "Добавить точку", onClick: ()=>{
                setSelectedBranch(branch)
                setMopen("Новая точка")
            }})
        return {
            key: branch.ID,
            icon: <FontAwesomeIcon icon={faShop} />,
            children: sps,
            label: branch.name
        }
    }

    const createBranch = (data) => {
        axios.post(wapi("adm/branch"), data, {headers: {"Authorization": props.token}})
            .then(r => {
                toast.success("Заведение создано")
                mutate(wapi("adm/branches"))
                setMopen(false)
            }).catch(e=>{
                toast.error("Не удалось создать заведение")
        })
    }

    const createSP = (data) => {
        axios.post(wapi("adm/sp"), data, {headers: {"Authorization": props.token}})
            .then(r => {
                toast.success("Точка создана")
                mutate(wapi("adm/branches"))
                setMopen(false)
            }).catch(e=>{
                toast.error("Не удалось создать точку")
        })
    }

    return <div className="flex flex-col gap-2">
        <p className="flex items-center justify-between">
            <span className="text-xl">Точки обслуживания</span>
            <Button type="primary" className="bg-blue-600 w-fit ml-auto" onClick={()=>setMopen("Новое заведение")}>
                <FontAwesomeIcon className="mr-2" icon={faAdd}/> Добавить
            </Button>
        </p>
        <Menu className="rounded-lg" mode="inline" items={data?.map(branchView)} />
        <Modal open={mopen} title={mopen} onCancel={()=>setMopen(false)} footer={[]}>
            {mopen==="Новое заведение" &&
                <Form name="basic"
                      labelCol={{span: 8}}
                      wrapperCol={{span: 16}}
                      initialValues={{remember: true}}
                      onFinish={createBranch}
                      autoComplete="off">
                    <Form.Item
                        label="Название"
                        name="name"
                        rules={[{
                                required: true,
                                message: 'Введите название',
                            }]}>
                        <Input />
                    </Form.Item>
                    <Form.Item className="flex justify-end">
                        <Button type="primary" htmlType="submit" className="bg-blue-600">
                            Создать
                        </Button>
                    </Form.Item>
                </Form>}
            {mopen==="Новая точка" &&
                <Form name="basic"
                      labelCol={{span: 8}}
                      wrapperCol={{span: 16}}
                      initialValues={{remember: true}}
                      onFinish={createSP}
                      autoComplete="off">
                    <Form.Item
                        label="Заведение"
                        name="branch" initialValue={selectedBranch.ID}>
                        <p>{selectedBranch.name}</p>
                    </Form.Item>
                    <Form.Item
                        label="Метки"
                        name="labels">
                        <Select mode="multiple" placeholder="Выберите метки" options={labels.map(lbl=>({label: lbl.name, value: lbl.ID}))} />
                    </Form.Item>
                    <Form.Item className="flex justify-end">
                        <Button type="primary" htmlType="submit" className="bg-blue-600">
                            Создать
                        </Button>
                    </Form.Item>
                </Form>}
        </Modal>
    </div>
}
