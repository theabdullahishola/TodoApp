import { FaPlus } from "react-icons/fa"

export default function Form (prop){


   
return(
    <>
    <form onSubmit={prop.handleForm}>
        <input name="todo" type="text"    />
        <button className="plus-button">
             <FaPlus />
        </button>
    </form>
    
    </>
)
}