
import { IoCheckmarkCircle } from "react-icons/io5";
import { MdDeleteForever } from "react-icons/md";

export default function TodoItem({ todo, deleteTodo, updateTodo }) {

  return (
    <li className={`list ${todo.completed ? "checked" : "unchecked"}`}>
      <span>{todo.body}</span>

      <button className="checkmark" onClick= {updateTodo}>
        <IoCheckmarkCircle />
      </button>

      <button className="delete" onClick={deleteTodo}>
        <MdDeleteForever />
      </button>
    </li>
  )
}