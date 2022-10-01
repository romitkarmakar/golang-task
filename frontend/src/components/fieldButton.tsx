import React, { useRef } from "react";

function FieldButton(props:{title:string, send:(arg:string)=>void}) {

    const inputRef = useRef<HTMLInputElement>(null)

    function handleClick() {
        console.log("Sending")
        if(!inputRef.current) return
        const reqBody = JSON.stringify({[`${props.title}`]:inputRef.current.value})
        props.send(reqBody)
        inputRef.current.value = ""
    }

    return <div>
        <input placeholder={props.title} ref={inputRef}></input>
        <button onClick={handleClick}> Send </button>
        </div>
}

export {FieldButton}