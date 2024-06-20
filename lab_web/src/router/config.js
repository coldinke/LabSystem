import Home from "../views/home/Home.vue";
import UserList from "../views/userManager/UserList.vue";
import RightList from "../views/rightManger/RightList.vue";
import RoleList from "../views/rightManger/RoleList.vue";
import LabList from "../views/labManager/LabList.vue";
import BookList from "../views/bookManger/BookList.vue";
import AuditList from "../views/bookManger/AuditList.vue";
import BookLab from "../views/bookManger/BookLab.vue";


const routes = [
  {
    path: "/index",
    component: Home,
  },
  {
    path: "/user-manage/userList",
    component: UserList,
  },
  {
    path: "/right-manage/roleList",
    component: RoleList,
  },
  {
    path: "/right-manage/rightList",
    component: RightList,
  },
  {
    path: "/lab-manage/labList",
    component: LabList,
  },
  {
    path: "/book-manage/bookList",
    component: BookList,
  },
  {
    path: "/book-manage/auditList",
    component: AuditList,
  },
  {
    path: "/book-manage/bookLab",
    component: BookLab
  }
];

export default routes;
