import { useRouter } from 'next/navigation'

export default function Input() {
  const router = useRouter()
  async function submit() {
    const data = {
      name: document.querySelector("input[name=name]").value,
      email: document.querySelector("input[name=email]").value,
      password: document.querySelector("input[name=password]").value
    }
    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/user/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    router.push('/')
  }
  return (
    <>
      <h1>User Create</h1>
      <div>
      <label>name: </label>
      <input type={"text"} name={"name"} required></input>
      </div>
      <div>
      <label>email: </label>
      <input type={"text"} name={"email"} required></input>
      </div>
      <div>
      <label>password: </label>
      <input type={"password"} name={"password"} required></input>
      </div>
      <div>
      <button onClick={submit}>add</button>
      </div>
    </>
  )
}
