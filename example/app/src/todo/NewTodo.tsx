import { useNavigate } from "react-router-dom";
import { useCreateTodoMutation } from "../gen/todo.api";
import TodoForm, { OnSubmit } from "./TodoForm";

export default function NewTodo() {

  const [createTodo] = useCreateTodoMutation()

  const navigate = useNavigate();

  const onSubit: OnSubmit = (data, isDirty) => {
    createTodo(data)
    navigate("/")
  }

  return <TodoForm onSubmit={onSubit} />

}
