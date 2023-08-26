
import Link from 'next/link'
import { useRouter } from 'next/navigation'

export default function Home({ taskList }) {
  const router = useRouter()
  async function deleteTask(id) {
    await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/delete/" + id, {
      method: "POST",
    })
    router.push('/')
  }
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

export async function getStaticProps() {
  const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/task")
  const taskList = await res.json()
  return {
    props: {
      taskList
    }
  }
}
