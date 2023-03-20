import * as React from "react";
import { useForm, useFormState } from "react-hook-form";
import { FieldMask } from "../gen/google/protobuf/field_mask";
import { State, Todo } from "../gen/todo";

interface TodoFormProps {
  value?: Todo
}

export default function TodoForm(props: TodoFormProps) {

  const { register, handleSubmit, control } = useForm<Todo>({
    defaultValues: props.value
  });

  const { isDirty, dirtyFields } = useFormState<Todo>({
    control
  });

  const onSubmit = (data: Todo) => {
    console.log(isDirty)
    console.log(toFieldMask(dirtyFields))
    console.log(data)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input type="hidden" {...register("id")} />
      <input {...register("text")} placeholder="Buy bitcoin" />
      <select {...register("state")} >
        {Object.entries(State).filter((a => !isNaN(Number(a[0])))).map((e) => {
          const [k, v] = e
          return <option key={k} value={k}>{v}</option>
        })}
      </select>
      <input type="number" {...register("loc.longitude")} />
      <input type="number" {...register("loc.latitude")} />
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

