import { Todo } from "./gen/todo";
import { useListTodosQuery } from "./gen/todo.api";
import './App.css';

interface TodoListProps {
}

export default function TodoList(props: TodoListProps) {
  const {
    data: todos,
    isFetching,
    isLoading,
  } = useListTodosQuery({})

  if (isLoading) return <div>Loading...</div>
  if (!todos) return <div>Missing todos!</div>

  const onEdit = (id: string) => {}
  const onDelete = (id: string) => {}

  return (
    <ul>
      {todos.items.map((item: Todo) => (
        <li key={item.id}>
          {item.title}
          <button onClick={() => onEdit(item.id)}>Edit</button>
          <button onClick={() => onDelete(item.id)}>Delete</button>
        </li>
      ))}
    </ul>
  );
    
}
