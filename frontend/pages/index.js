
export default function Home({ taskList }) {
  return (
    <>
      <h1>Task</h1>
      <ul>
      {taskList.map(({ Id, Name, CreatedAt }) => (
        <li key={Id}>{Name}({CreatedAt})</li>
      ))}
      </ul>
    </>
  )
}

export async function getStaticProps() {
  const res = await fetch(process.env.API_HOST + "/api/task")
  const taskList = await res.json()
  return {
    props: {
      taskList
    }
  }
}
