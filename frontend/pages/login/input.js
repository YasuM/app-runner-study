
export default function Input() {
  async function submit () {
    const data = {
      email : document.querySelector("input[name=email]").value,
      password: document.querySelector("input[name=password]").value
    }
    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
      credentials: "include"
    })
    // router.push('/')
  }
  return (
    <>
      <h1>Login</h1>
      <div>
      <label>email: </label>
      <input type={"text"} name={"email"} required></input>
      </div>
      <div>
      <label>password: </label>
      <input type={"password"} name={"password"} required></input>
      </div>
      <div>
      <button onClick={submit}>edit</button>
      </div>
    </>
  )
}
