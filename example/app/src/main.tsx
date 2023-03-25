import React from 'react'
import ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'
import App from './App'
import './index.css'
import { store } from './store'
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { grpcWebOptions } from "./gen/todo.api"
import TodoList from './todo/TodoList'
import EditTodo from './todo/EditTodo'
import NewTodo from './todo/NewTodo'

grpcWebOptions.baseUrl = "http://localhost:5080"

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        index: true,
        element: <TodoList />,
      },
      {
        path: "/todo/:todoId",
        element: <EditTodo />,
        loader: async ({ params }) => {
          return params.todoId
        },
      },
      {
        path: "/new",
        element: <NewTodo />,
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Provider store={store}>
      <RouterProvider router={router} />
    </Provider>
  </React.StrictMode>,
)
