
import Navbar from '../components/Navbar'
import Form from '../components/Form'
import FormList from '../components/FormList'
import { useQuery, useMutation,useQueryClient } from '@tanstack/react-query'
import { addTodo,getTodos,deleteTodo,updateTodo } from '../api/todo'
import LogOut from '../components/Logout'


export default function TodoDashboard() {
  const queryClient = useQueryClient()
 const {isPending, isError, data, error}= useQuery({queryKey: ["todos"],
  queryFn:getTodos}) 

 const mutation = useMutation({
  mutationFn: addTodo,
  onSuccess: () => {
    // Invalidate and refetch
    queryClient.invalidateQueries({ queryKey: ['todos'] })
  },
})
const updateMutation = useMutation({
  mutationFn: updateTodo,
  onSuccess: () => {
    // Invalidate and refetch
    queryClient.invalidateQueries({ queryKey: ['todos'] })
  },
})
const deleteMutation = useMutation({
  mutationFn: deleteTodo,
  onSuccess: () => {
    // Invalidate and refetch
    queryClient.invalidateQueries({ queryKey: ['todos'] })
  },
})


function handleForm(e){
    e.preventDefault()
    const form= e.target
    const newTodo = form.todo.value
    if (!newTodo.trim()) return // ignore empty input
    mutation.mutate(newTodo)
    form.reset()
  }
 
function updateTodoHandler ({id, updatedTodo}){
  updateMutation.mutate({id, updatedTodo})
}
   
  const deleteHander = (id) => {
    deleteMutation.mutate(id)
  }

  return (
    <>
      <Navbar />
      <LogOut />
      {isPending && <p>Loading...</p>}
      {isError && <p>Error: {error.message}</p>}

      {!isPending && !isError && (
        <>
          <Form handleForm={handleForm} />
          <FormList todos={data} deleteTodo={deleteHander} updateTodo={updateTodoHandler}/>
        </>
      )}
    </>
  )
}


