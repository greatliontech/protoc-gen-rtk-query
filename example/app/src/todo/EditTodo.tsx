import { useLoaderData } from "react-router-dom";
import { useGetTodoQuery } from "../gen/todo.api";
import TodoForm from "./TodoForm";

export default function EditTodo() {

  const id = useLoaderData() as string;
  const {data, isLoading} = useGetTodoQuery({ id })

  if (isLoading) return <div>Loading...</div>

  return <TodoForm value={data} />
}

