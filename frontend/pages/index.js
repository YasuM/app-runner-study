
import Link from 'next/link'

export default function Home({ taskList }) {
  return (
    <>
      <h1>Task</h1>
      <ul>
      {taskList.map(({ Id, Name, CreatedAt }) => (
        <li key={Id}>{Name}({CreatedAt})</li>
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
