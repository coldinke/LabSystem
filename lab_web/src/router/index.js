import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../views/Login.vue'
import MainBox from '../views/MainBox.vue'
import LabDetail from '../components/lab/LabDetail.vue'
import NotFound from '../views/notfound/NotFound.vue'
import RouteConfig from './config'
import { useRouterStore } from '../store/useRouterStore'
import { useUserStore } from '../store/useUserStore'

const routes = [
    {
        path: "/login",
        name: "login",
        component: Login
    },
    {
        path: "/mainbox",
        name: "mainbox",
        component: MainBox
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

router.beforeEach((to, from, next) => {
    const { isGetterRouter } = useRouterStore()
    const { user } = useUserStore()
    if (to.name === "login") {
        next()
    } else {
        if (!user.role) {
            next({
                path: "/login"
            })
        } else {
            if (!isGetterRouter) {
                router.removeRoute("mainbox")
                ConfigRouter()
                next({
                    path: to.fullPath
                })
            } else {
                next()
            }
        }
    }

})

// add router for auth.
const ConfigRouter = () => {
    let { changeRouter } = useRouterStore()

    router.addRoute({
        path:"/mainbox",
        name:"mainbox",
        component: MainBox,
    })

    RouteConfig.forEach(item => {
        checkPermission(item.path) && router.addRoute("mainbox", item)
    })

    router.addRoute("mainbox", {
        path: "/",
        redirect: "/index"
    })

    router.addRoute("lab-detail", {
        path: "/labs/:labId", // 这里的 :labId 是参数占位符
        component: LabDetail,
      });
    
    router.addRoute("mainbox",{ 
        path: '/:pathMatch(.*)*', 
        name: 'NotFound', 
        component: NotFound 
    })

    // console.log(router.getRoutes())
    changeRouter(true)
}

const checkPermission = (path=>{
    const {user} = useUserStore()

    return user.role.rights.some(right => right.path === path);
})

export default router