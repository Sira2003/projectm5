import React, { useState } from "react";
import { UserOutlined, 
  
  TeamOutlined ,
  MenuOutlined, 
  HomeOutlined, 
  ScheduleOutlined, 
  CreditCardOutlined,
  MedicineBoxOutlined,
  IdcardOutlined,
  AuditOutlined,
  PlusOutlined } from "@ant-design/icons";

import type { MenuProps } from "antd";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
} from "react-router-dom";

import { Breadcrumb, Layout, Menu, theme } from "antd";



// import Dashboard from "./pages/dashboard";
import Employee from "./pages/employee/index.tsx";
import Patient from "./pages/patient/index.tsx";
import PatientCreate from "./pages/patient/create";
import EmployeeCreate from "./pages/employee/create";
import EmployeeEdit from "./pages/employee/edit";
 import PatientEdit from "./pages/patient/edit";

import Schedule from "./pages/schedule/create.tsx";
import Schedule2 from "./pages/schedule/view.tsx";


const { Header, Content, Footer, Sider   } = Layout;
const { SubMenu } = Menu;

type MenuItem = Required<MenuProps>["items"][number];

// function getItem(
//   label: React.ReactNode,
//   key: React.Key,
//   icon?: React.ReactNode,
//   children?: MenuItem[]
// ): MenuItem {
//   return {
//     key,
//     icon,
//     children,
//     label,
//   } as MenuItem;
// }


const App: React.FC = () => {
  // const page = localStorage.getItem("page");
  // const [collapsed, setCollapsed] = useState(false);
  // const {
  //   token: { colorBgContainer },
  // } = theme.useToken();
  const [selectedKeys, setSelectedKeys] = useState<string[]>(["dashboard"]);
  const [openKeys, setOpenKeys] = useState<string[]>([]);
  const [collapsed, setCollapsed] = useState(false);

  // const setCurrentPage = (val: string) => {
  //   localStorage.setItem("page", val);
  // };
  const onOpenChange = (keys: string[]) => {
    setOpenKeys(keys.length === 0 ? [] : [keys[keys.length - 1]]);
  };

  const onMenuClick = (e: { key: string }) => {
    setSelectedKeys([e.key]);
  };
  
  // const handleMenuClick = (e: { key: string }) => {
  //   setSelectedKey(e.key);
  //   localStorage.setItem("page", e.key);
  // };

  return (
    <Router>
      <Layout style={{ minHeight: "100vh" }}>
        <Sider
          collapsible
          collapsed={collapsed}
          onCollapse={(value) => setCollapsed(value)}
          style={{
            background: "linear-gradient(180deg, #22225E 0%, #22225E 20%, #7DC9D1 80%, #42C2C2 100%)"
          }}
        >
          
          <div
            style={{
              display: "flex",
              justifyContent: "center",
              marginTop: 20,
              marginBottom: 150,
            }}
          >
            
            {/* <img
              src={logo}
              alt="Logo"
              style={{ width: "80%"}}
            /> */}
            <div 
              style={{display: "flex", alignItems: "center", color: "white",}}
            >
              <MenuOutlined style= {{color: "white",marginRight: 8,}} />
              <div >SAMGT PLEASANT</div>
              
            </div>
          </div>
          <Menu
             theme="dark"
             mode="inline"
             selectedKeys={selectedKeys} 
             openKeys={openKeys}         
             onOpenChange={onOpenChange} // ฟังก์ชันจัดการการเปิด/ปิด SubMenu
             onClick={onMenuClick}       // ฟังก์ชันจัดการการเลือกเมนู
             style={{
               background: "transparent",
             }}
          >
            {/* <Menu.Item key="dashboard" className="d-md-none" onClick={() => setCurrentPage("dashboard")}> */}
            <Menu.Item key="dashboard" >
              <Link to="/">
                {/* <DashboardOutlined /> */}
                {/* <img
                src={home}
                alt="home"
                style={{ maxWidth: "150px", height: "auto" , marginRight: 6,}}
                /> */}
                <HomeOutlined />
                <span >หน้าหลัก</span>
              </Link>
            </Menu.Item>
            {/* // อย่าลืมแก้ key */}

            {/* <Menu.Item key="customer" className="d-md-none" onClick={() => setCurrentPage("customer")}>
              <Link to="/customer"> 
              <IdcardOutlined />
                <span>จัดการประวัติ</span>
              </Link>
            </Menu.Item> */}

            <SubMenu 
            key="manageRecords" 
            
            title={<span><IdcardOutlined />
            <span>จัดการประวัติ</span></span>}>
              {/* <Menu
              //theme="dark"
              style={{ backgroundColor: 'rgba(66, 194, 194, 0.5)',marginLeft: 8, }}> */}
                <Menu.Item key="managePatients" style={{
                  backgroundColor:
                    selectedKeys.includes("managePatients") ? "#1890ff" : "rgba(66, 194, 194, 0.5)", // เปลี่ยนสีเมื่อถูกเลือก
                  marginLeft: 8,
                }}  >
                  <Link to="/patients">
                  <UserOutlined />
                  <span>ผู้ป่วย</span>
                  </Link>
                </Menu.Item>
                <Menu.Item key="manageEmployees" style={{
                  backgroundColor:
                    selectedKeys.includes("manageEmployees") ? "#1890ff" : "rgba(66, 194, 194, 0.5)", // เปลี่ยนสีเมื่อถูกเลือก
                  marginLeft: 8,
                }} >
                  <Link to="/employees">
                  <TeamOutlined />
                  <span>พนักงาน</span>
                  </Link>
                </Menu.Item>
              {/* </Menu> */}
            </SubMenu>

            <Menu.Item key="cussstomer" >
              <Link to="/schedule">
              <ScheduleOutlined />
              <span>นัดหมาย</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="cusssstomer"  >
              <Link to="/schedule2">
              <CreditCardOutlined />
              <span>การเงิน</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="cussssstomer"  >
              <Link to="/customer">
              <MedicineBoxOutlined />
              <span>อุปกรณ์</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="cussssstomerr"  >
              <Link to="/customer">
              <AuditOutlined />
              <span>รายการเบิก</span>
              </Link>
            </Menu.Item>
            <Menu.Item key="cussssstomerrr"  >
              <Link to="/customer/create">
              <PlusOutlined />
              <span>บันทึกการรักษา</span>
              </Link>
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout>
          {/* <Header style={{ padding: 0, background: colorBgContainer }} /> */}
          <Header style={{ padding: 0, background: theme.useToken().token.colorBgContainer }} />
          <Content style={{ margin: "0 16px" }}>
            {/* <Breadcrumb style={{ margin: "16px 0" }} /> */}
            <div
              style={{
                padding: 24,
                minHeight: "100%",
                // background: colorBgContainer,
                background:theme.useToken().token.colorBgContainer
              }}
            >
              <Routes>
                {/* <Route path="/" element={<Dashboard />} />
                <Route path="/customer" element={<Customer />} />
                <Route path="/customer/create" element={<CustomerCreate />} />
                <Route path="/customer/edit/:id" element={<CustomerEdit />} /> */}
                <Route path="/" element={<div>Manage Patient Records</div>} />
                <Route path="/customer" element={<div>Manage Patient Records</div>} />
                <Route path="/employee/create" element={<EmployeeCreate/>} />
                <Route path="/employee/edit/:id" element={<EmployeeEdit/>} />
                <Route path="/manage-patients" element={<div>Manage Patient Records</div>} />
                <Route path="/employees" element={<Employee />} />
                <Route path="/schedule" element={<Schedule />} />
                <Route path="/schedule2" element={<Schedule2 />} />
                
                <Route path="/patients" element={<Patient/>} />
                <Route path="/patient/create" element={<PatientCreate/>} />
                <Route path="/patient/edit/:id" element={<PatientEdit/>} />
              </Routes>
            </div>
          </Content>
          {/* <Footer style={{ textAlign: "center" }}>
            System Analysis and Design 1/67
          </Footer> */}
        </Layout>
      </Layout>
    </Router>
  );
};

export default App;
