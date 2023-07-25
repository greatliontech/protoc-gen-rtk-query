import { Todo } from "@buf/greatliontech_protoc-gen-rtk-query-example.community_timostamm-protobuf-ts/todo_pb";
import { useDeleteTodoMutation, useListTodosQuery } from "../gen/todo.api";
import { useNavigate } from "react-router-dom";

import './TodoList.css';

export default function TodoList() {

  const { data: todos, isLoading } = useListTodosQuery({})

  const [deleteTodo] = useDeleteTodoMutation()

  const navigate = useNavigate();

  if (isLoading) return <div>Loading...</div>
  if (!todos) return <div>Missing todos!</div>

  const onEdit = (id: string) => {
    navigate(`/todo/${id}`)
  }

  const onDelete = (id: string) => {
    deleteTodo({ id })
  }

  return (
    <ul>
      {todos.items.map((item: Todo) => (
        <li key={item.id}>
          <span>{item.title}</span>
          <button onClick={() => onEdit(item.id)}>E</button>
          <button onClick={() => onDelete(item.id)}>D</button>
        </li>
      ))}
    </ul>
  );

}
