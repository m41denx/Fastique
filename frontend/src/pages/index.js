import {Rubik} from "next/font/google"
import Link from "next/link";

const rubik = Rubik({subsets: ["cyrillic"]})


export default function Index(props) {
  return <div className="h-screen flex flex-col justify-center items-center bg-hero">
    <p className="text-2xl">Максимально дефолтные</p>
    <p className={`text-7xl ${rubik.className}`}>Билетики</p>
    <p className="tracking-[1rem] bg-gradient-radial from-red-600  to-transparent px-8 py-1 rounded-xl">В Бобруйск</p>
    <div className="flex gap-4 mt-4">
      <Link href="/order" className="px-4 py-2 bg-blue-600 hover:bg-blue-800 cursor-pointer text-white text-2xl rounded-xl">Получить талон</Link>
      <Link href="/process" className="px-4 py-2 bg-green-600 hover:bg-green-800 cursor-pointer text-white text-2xl rounded-xl">Я сотрудник</Link>
    </div>
  </div>
}