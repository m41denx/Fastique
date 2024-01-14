import {Button, Form, Input} from "antd";
import toast, {Toaster} from "react-hot-toast";
import axios from "axios";
import wapi from "@/config";
import {useRouter} from "next/router";


export default function Login(props) {

    const router = useRouter()

    const onFinish = async (v)=>{
        let data = await axios.post(wapi("auth/login"), v, {validateStatus: () => true}).then(data=>{
            if (data.status === 200){
                localStorage.setItem("token", data.data.token)
                props.setToken(data.data.token)
                router.push("/process")
            } else {
                toast.error("Неверное имя пользователя или пароль")
            }
        }).catch(e=>{
            console.log(e)
            toast.error("Ошибка соединения")
        })

    }

    return <div className="h-screen flex flex-col justify-center items-center">
        <Toaster />
        <Form name="basic"
            labelCol={{
                span: 8,
            }}
            wrapperCol={{
                span: 16,
            }}
            style={{
                maxWidth: 600,
            }}
            initialValues={{
                remember: true,
            }}
              className=" rounded-xl p-4 shadow drop-shadow"
            onFinish={onFinish}
            autoComplete="off"
        >
            <p className="text-lg text-center mb-4">Вход для сотрудников</p>
            <Form.Item
                label="Логин"
                name="username"
                rules={[
                    {
                        required: true,
                        message: 'Введите имя пользователя',
                    },
                ]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="Пароль"
                name="password"
                rules={[
                    {
                        required: true,
                        message: 'Введите пароль',
                    },
                ]}
            >
                <Input.Password />
            </Form.Item>

            <Form.Item
                wrapperCol={{
                    offset: 8,
                    span: 16,
                }}
            >
                <Button type="primary" htmlType="submit" className="bg-blue-600">
                    Войти
                </Button>
            </Form.Item>
        </Form>
    </div>
}