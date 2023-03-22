import React from 'react'
import ReactDOM from 'react-dom/client'
import { Provider } from 'react-redux'
import App from './App'
import './index.css'
import { store } from './store'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import { grpcWebOptions } from "./gen/todo.api"
import TodoForm from './TodoForm'

grpcWebOptions.baseUrl = "http://localhost:5080"

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/todo:todoId",
        element: <TodoForm />,
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
