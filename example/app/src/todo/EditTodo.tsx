import { FieldMask } from "../gen/google/protobuf/field_mask_pb";
import { useLoaderData, useNavigate } from "react-router-dom";
import { useGetTodoQuery, useUpdateTodoMutation } from "../gen/todo.api";
import TodoForm, { OnSubmit } from "./TodoForm";

export default function EditTodo() {

  const id = useLoaderData() as string;
  const { data, isLoading } = useGetTodoQuery({ id })

  const [updateTodo] = useUpdateTodoMutation()

  const navigate = useNavigate();

  const onSubit: OnSubmit = (formData, dirtyFields) => {
    console.log("onSubmit", formData, dirtyFields, data)
    updateTodo({
      todo: formData,
      updateMask: toFieldMask(dirtyFields),
    })
    navigate("/")
  }

  if (isLoading) return <div>Loading...</div>

  return <TodoForm value={data} onSubmit={onSubit} />
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
