import { useForm, useFormState } from "react-hook-form";
import { FieldMask } from "./gen/google/protobuf/field_mask";
import { State, Todo } from "./gen/todo";
import { useCreateTodoMutation, useUpdateTodoMutation } from "./gen/todo.api";

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

  const [createTodo, createResult] = useCreateTodoMutation()
  const [updateTodo, updateResult] = useUpdateTodoMutation()

  const onSubmit = (data: Todo) => {
    console.log("on submit. is dirty?:", isDirty)
    console.log(toFieldMask(dirtyFields))
    console.log(data)
    if (props.value) {
      console.log("updating")
      updateTodo({
        todo: { 
          id: data.id,
          title: data.title,
          description: data.description,
          state: 0,
        },
        updateMask: toFieldMask(dirtyFields),
      })
      return
    }
    console.log("creating")
    createTodo({ 
      id: data.id,
      title: data.title,
      description: data.description,
      state: 0,
    },)
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input type="hidden" {...register("id")} />
      <input {...register("title")} placeholder="Buy bitcoin" />
      <input {...register("description")} />
      <select {...register("state")} >
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

