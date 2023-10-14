import { useRouter } from 'next/navigation'

export default function Input() {
  const router = useRouter()

  async function submit() {
    const data = {task: document.querySelector("input[name=task]").value}
    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: 'include',
      mode: 'cors',
      body: JSON.stringify(data),
    });
    router.push('/')
  }
  return (
    <>
      <h1>Task Input</h1>
      <label>task: </label>
      <input type={"text"} name={"task"} required></input>
      <div>
      <button onClick={submit}>aaa</button>
      </div>
    </>
  )
}
