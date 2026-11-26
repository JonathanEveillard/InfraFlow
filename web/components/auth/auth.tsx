import {useState} from "react";
import "./auth.css"

export default function Auth({isAuth}:{isAuth:boolean}){
    const [form, setForm] = useState({
        username:"",
        password: ""
    })

    function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
        const {name,value} = e.target;
        setForm({...form, [name]:value});
    }

    function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        console.log(form)
    }

    return isAuth?(
        <div>
            TO IMPLEMENT AKA REDIRECT
        </div>
    ) : (
        <div className="auth-container">
            <form className="auth-form" onSubmit={handleSubmit}>
                <label htmlFor="username">Username</label>
                <input
                    id="username"
                    name="username"
                    value={form.username}
                    onChange={handleChange}
            />

                <label htmlFor="password">Password</label>
                <input
                    id="password"
                    name="password"
                    type="password"
                    value={form.password}
                    onChange={handleChange}
                />
                <button type="submit">Log in</button>
        </form>
    </div>
    )
}