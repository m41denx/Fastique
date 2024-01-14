import {Rubik} from "next/font/google"
import Link from "next/link";

const rubik = Rubik({subsets: ["cyrillic", "latin"]})


export default function Index(props) {
  return <div className="h-screen flex flex-col justify-center items-center bg-hero">
    <p className="text-2xl">Максимально дефолтные</p>
    <p className={`text-7xl ${rubik.className}`}>Билетики</p>
    <div className="grid grid-cols-2 gap-4 mt-4 rounded-xl bg-white bg-opacity-50 backdrop-blur shadow-lg drop-shadow-2xl p-2">
      <div className="flex flex-col gap-2 items-center">
        <p className="text-xl">Зона пользователя</p>
        <Link href="/order" className="text-center px-4 py-2 w-full bg-blue-600 hover:bg-blue-800 cursor-pointer text-white text-lg rounded-xl">Запись онлайн</Link>
        <Link href="/orderlocal" className="text-center px-4 py-2 w-full bg-blue-600 hover:bg-blue-800 cursor-pointer text-white text-lg rounded-xl">Запись на месте</Link>
      </div>
      <div className="flex flex-col gap-2 items-center">
        <p className="text-xl">Зона компании</p>
        <Link href="/process" className="text-center px-4 py-2 w-full bg-green-600 hover:bg-green-800 cursor-pointer text-white text-lg rounded-xl">Место сотрудника</Link>
        <Link href="/screen" className="text-center px-4 py-2 w-full bg-green-600 hover:bg-green-800 cursor-pointer text-white text-lg rounded-xl">Экран очереди</Link>
      </div>
    </div>
  </div>
}