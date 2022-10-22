package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"time"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello world, %s!", request.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time%s", time.Now().String())
}
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "find")
}

func handler4(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test")
}

func HeaderHander(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func ProcessHandler2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

// 测试ResponseWriter
func WriteTest(w http.ResponseWriter, r *http.Request) {
	str := `<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset = utf-8">
        <title>GoWeb开发</title>
    </head>
    <body>
		ResponserWriter测试
	</body>
	</html>`
	w.Write([]byte(str))
}

// 修改状态码
func WriteHeaderTest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "未知服务！")
}

type Post struct {
	Name    string
	Threads []string
}

// 响应返回json
func JsonTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		Name:    "急速小子",
		Threads: []string{"早", "中", "晚"},
	}
	jsonPost, _ := json.Marshal(post)
	w.Write(jsonPost)
}

// 向客户端发送cookie
func SetCookieTest(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "cookie_1",
		Value:    "小饼干",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "cookie_2",
		Value:    "small cookie",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
}

func GetCookieTest(w http.ResponseWriter, r *http.Request) {
	// cookies := r.Header["Cookie"]
	// cookies := r.Cookies()
	cookie, err := r.Cookie("cookie_2")
	if err != nil {
		fmt.Fprintln(w, "没有此cookie")
	}
	fmt.Fprintln(w, cookie)
}

// 使用模板引擎
func TempalteTest(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./first_webapp/src/html/tmpl.html")
	if err != nil {
		fmt.Println("出错了")
	}
	t.Execute(w, "hello world!")
}

// 模板-包含动作
func TemplateBaoHanTest(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./first_webapp/src/html/t1.html", "./first_webapp/src/html/t2.html"))
	t.Execute(w, "hello")
}

// XSS攻击
func XSStest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./first_webapp/src/html/tmpl.html")
	t.Execute(w, r.FormValue("comment"))
}
func FormTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./first_webapp/src/html/form.html")
	t.Execute(w, nil)
}
func XSSNotTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./first_webapp/src/html/tmpl.html")
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

// REST风格获取URL中的参数,path.Base()
func RESTTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Println(path.Base("/a/b/c/d.abc"))
	id, err := strconv.Atoi(path.Base((r.URL.Path)))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id是", id)
	fmt.Fprint(w, id)
	return
}

func main() {
	http.HandleFunc("/", handler)       //8080/ 8080 ;http://localhost:8080/find/time
	http.HandleFunc("/time", handler2)  //8080/time; http://localhost:8080/time
	http.HandleFunc("/find", handler3)  // http://localhost:8080/find
	http.HandleFunc("/test/", handler4) // http://localhost:8080/test/;http://localhost:8080/test;http://localhost:8080/test/find
	http.HandleFunc("/process", ProcessHandler)
	http.HandleFunc("/postform", ProcessHandler2)
	http.HandleFunc("/write", WriteTest)
	http.HandleFunc("/writeHeader", WriteHeaderTest)
	http.HandleFunc("/json", JsonTest)
	http.HandleFunc("/setcookie", SetCookieTest)

	http.HandleFunc("/getcookie", GetCookieTest)
	http.HandleFunc("/template", TempalteTest)
	http.HandleFunc("/template-baohan", TemplateBaoHanTest)

	http.HandleFunc("/xss", XSStest)
	http.HandleFunc("/xssnot", XSSNotTest)
	http.HandleFunc("/form", FormTest)

	http.HandleFunc("/post/", RESTTest)

	//header测试
	http.HandleFunc("/header", HeaderHander)
	http.ListenAndServe(":8020", nil)

}
