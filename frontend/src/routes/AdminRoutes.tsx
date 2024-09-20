import { lazy } from "react";
import { RouteObject } from "react-router-dom";
import Loadable from "../component/third-party/Loadable";
import FullLayout from "../layout/FullLayout";


const MainPages = Loadable(lazy(() => import("../pages/authentication/Login")));
//const Dashboard = Loadable(lazy(() => import("../pages/dashboard")));
const Employee = Loadable(lazy(() => import("../pages/employee")));
const CreateEmployee = Loadable(lazy(() => import("../pages/employee/create")));
const EditEmployee = Loadable(lazy(() => import("../pages/employee/edit")));

const Patient = Loadable(lazy(() => import("../pages/patient")));
const CreatePatient = Loadable(lazy(() => import("../pages/patient/create")));
const EditPatient = Loadable(lazy(() => import("../pages/patient/edit")));


const AdminRoutes = (isLoggedIn : boolean): RouteObject => {
  return {
    path: "/",
    element: isLoggedIn ? <FullLayout /> : <MainPages />,
    children: [
      {
        path: "/",
        element: <MainPages />,
        //element: <div>test</div>,

      },
      {
        path: "/patient",
        children: [
          {
            path: "",
            element: <Patient />,
          },
          {
            path: "create",
            element: <CreatePatient />,
          },
          {
            path: "edit/:id",
            element: <EditPatient />,
          },
        ],
      },
      {
        path: "/employee",
        children: [
          {
            path: "employee",
            element: <Employee/>,
          },
          {
            path: "create",
            element: <CreateEmployee />,
          },
          {
            path: "edit/:id",
            element: <EditEmployee />,
          },
        ],
      },
    ],
  };
};

export default AdminRoutes