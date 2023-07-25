import { FieldNamesMarkedBoolean, useForm, useFormState } from "react-hook-form";
import { State, Todo } from "@buf/greatliontech_protoc-gen-rtk-query-example.community_timostamm-protobuf-ts/todo_pb";

import './TodoForm.css';

export type OnSubmit = (data: Todo, dirtyFields: Partial<Readonly<FieldNamesMarkedBoolean<Todo>>>) => void

interface TodoFormProps {
  value?: Todo
  onSubmit?: OnSubmit
}

export default function TodoForm(props: TodoFormProps) {

  const { register, handleSubmit, control } = useForm<Todo>({
    defaultValues: props.value,
  });

  const { dirtyFields } = useFormState<Todo>({
    control
  });

  const onSubmit = (data: Todo) => {
    props.onSubmit?.(data, dirtyFields)
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
