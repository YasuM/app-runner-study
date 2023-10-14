import React, { useState, useEffect } from 'react'
import Link from 'next/link'
 
import { useRouter } from 'next/navigation'
export default function Index() {
 
  const [taskList, setTaskList] = useState([])
  const router = useRouter()
  async function deleteTask(id) {
    await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/delete/" + id, {
      method: "POST",
    })
    router.push('/')
  }

  useEffect(() => {
    const fetchData = async () => {
      console.log(process.env.NEXT_PUBLIC_API_HOST)
      const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/task", {
        method: 'GET',
        mode: 'cors',
        credentials: 'include'
      })
      const result = await res.json()
      setTaskList(result)
    }
 
    fetchData().catch((e) => {
      console.error('An error occurred while fetching the data: ', e)
    })
  }, [])

  return (
  <>
    <h1>Task</h1>
    <ul>
    {taskList.map(({ Id, Name, StatusLabel, CreatedAt }) => (
      <li key={Id}>{Name} {StatusLabel} {CreatedAt} <a href={`/edit?id=${Id}`}>edit</a> <button value={`${Id}`} onClick={e => deleteTask(e.target.value)}>delete</button></li>
    ))}
    </ul>
     <ul>
       <Link href="/add">Add</Link>
     </ul>
  </>
)
}
