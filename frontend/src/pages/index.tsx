import axios from "axios";
import { Inter } from "next/font/google";
import { useEffect, useState } from "react";
import configJson from "../config.json";


const inter = Inter({ subsets: ["latin"] });

export default function Home() {
  const [users, setUsers] = useState<any>([]);
  useEffect(() => {
    axios.get(`${configJson.localApi}/users`).then((res) => {
      setUsers(res.data)
    }).catch((err) => {
      alert("error happend")
    })
  },[])
  return (
    <main
      className={`flex min-h-screen flex-col items-center p-4 ${inter.className}`}
    >
      <h1 className="mb-3">All Users</h1>
      <div className="w-full grid grid-cols-1 gap-2 md:grid-cols-2">
        {users?.data?.items?.map((item: any) => {
          return (
            <li
              key={item.id}
              className="w-full mx-auto p-10 bg-sky-600 text-white rounded-md"
            >
              user email : {item.email}
            </li>
          );
        })}
      </div>
    </main>
  );
}
