import { useRouter } from 'next/navigation'

export default function Edit(params) {
  const taskStatusList = params.taskStatusList
  const router = useRouter()

  async function submit () {
    const data = {
      id: Number(params.task.Id),
      task: document.querySelector("input[name=task]").value,
      status: Number(document.querySelector("select[name=status]").value)
    }
    console.log(data)
    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/edit", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
    router.push('/')
  }
  return (
    <>
      <h1>Task Edit</h1>
      <div>
      <label>task: </label>
      <input type={"text"} name={"task"} defaultValue={params.task.Name} required></input>
      </div>
      <div>
      <label>status: </label>
      <select name="status">
          {taskStatusList.map((v, i) => <option key={i} value={v.Id}>{v.Label}</option>)}
      </select>
      </div>
      <div>
      <button onClick={submit}>edit</button>
      </div>
    </>
  )
}

export async function getServerSideProps({ query }) {
  const resTask = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/task/" + query.id)
  const resTaskStatus = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/task_status")
  const task = await resTask.json()
  const taskStatusList = await resTaskStatus.json()
  return {
    props: {
      taskStatusList,
      task,
    }
  }
}
