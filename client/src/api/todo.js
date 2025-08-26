
import api from "./axios";

//getTodo
export const getTodos = async () => {
    const res = await api.get("/api/todos/"); 
    return res.data; 
  };
  
  // Add a new todo
  export const addTodo = async (newTodo) => {
    const res = await api.post("/api/todos/", {body: newTodo});
    return res.data;
  };
  
  // Update a todo
  export const updateTodo = async ({ id, updatedTodo }) => {
    const res = await api.put(`/api/todos/${id}`,  updatedTodo);
    return res.data; 
  };
  
  // Delete a todo
  export const deleteTodo = async (id) => {
    const res = await api.delete(`/api/todos/${id}`);
    return res.data; // e.g., { message: "deleted" }
  };