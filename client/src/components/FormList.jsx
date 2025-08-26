import TodoItem from "./TodoItem"

export default function FormList({ todos, deleteTodo, updateTodo }) {
  if (!todos || todos.length === 0) {
    return null;  
  }

  return (
    <main>
      <ul>
        {todos.map((todo) => (
          <TodoItem 
            key={todo.id} 
            todo={todo} 
            deleteTodo={() => deleteTodo(todo.id)} 
            updateTodo={() => updateTodo({ id: todo.id, updatedTodo: { completed: !todo.completed } })} 
          />
        ))}
      </ul>
    </main>
  )
}
