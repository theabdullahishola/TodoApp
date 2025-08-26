import {logout} from '../api/auth'
import { AuthContext } from "../App";
import { useContext } from "react";

export default function LogOut() {
    const { setIsLoggingOut, isLoggingOut } = useContext(AuthContext);
function handleClick(){
    logout(setIsLoggingOut)
}

    return(
        <>
        <button onClick={handleClick}>LogOut</button>
        </>

    )

}