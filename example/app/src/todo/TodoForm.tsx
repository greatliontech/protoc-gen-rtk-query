import { useForm, useFormState } from "react-hook-form";
import { FieldMask } from "../gen/google/protobuf/field_mask";
import { State, Todo } from "../gen/todo";
import { useCreateTodoMutation, useUpdateTodoMutation } from "../gen/todo.api";
import { useNavigate } from "react-router-dom";

import './TodoForm.css';

interface TodoFormProps {
  value?: Todo
}

export default function TodoForm(props: TodoFormProps) {

  const { register, handleSubmit, control } = useForm<Todo>({
    defaultValues: props.value,
  });

  const { dirtyFields } = useFormState<Todo>({
    control
  });

  const navigate = useNavigate();

  const [createTodo] = useCreateTodoMutation()
  const [updateTodo] = useUpdateTodoMutation()

  const onSubmit = (data: Todo) => {
    if (props.value) {
      console.log("updating")
      updateTodo({
        todo: data,
        updateMask: toFieldMask(dirtyFields),
      })
    } else {
      console.log("creating")
      createTodo(data)
    }
    navigate("/")
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>

      <input type="hidden" {...register("id")} />

      <label htmlFor="title">Title</label>
      <input id="title" {...register("title")} placeholder="Buy bitcoin" />

      <label htmlFor="description">Description</label>
      <input id="description" {...register("description")} />

      <label htmlFor="state">Status</label>
      <select id="state" {...register("state", { setValueAs: value => Number(value) })} >
        {Object.entries(State).filter((a => !isNaN(Number(a[0])))).map((e) => {
          const [k, v] = e
          return <option key={k} value={k}>{v}</option>
        })}
      </select>
      <input type="submit" />
    </form>
  );
}

const toFieldMask = (o: Object): FieldMask => {

  const getFields = (o: Object): string[] => {
    let out: string[] = []
    Object.entries(o).map((e) => {
      let [k, v] = e
      if (typeof (v) === "object") {
        out = out.concat(getFields(v).map(s => k + "." + s))
      } else {
        out.push(k)
      }
    })
    return out
  }

  return FieldMask.fromJson(getFields(o).join(","))
}

