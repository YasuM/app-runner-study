
export default function Input() {
  async function submit() {
    const data = {task: document.querySelector("input[name=task]").value}
    const res = await fetch(process.env.NEXT_PUBLIC_API_HOST + "/api/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    console.log(res)
    
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