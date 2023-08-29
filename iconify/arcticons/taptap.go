package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taptap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.111 30.146c-4.517.804-10.837.416-13.803-7.733S10.04 8.125 17.755 7.03c8.593-1.22 18.349-.36 19.652 8.897c.806 5.728-.882 8.24-2.64 9.554"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.111 30.146c-.958.9-.776 12.583 12.085 11.53c12.86-1.053 10.81-10.81 9.202-12.972s-3.382-3.99-7.9-3.02c-4.517.97-11.557 2.743-13.387 4.462Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.122 38.345s-1.428 2.195-6.113 2.61S8.5 38.85 8.637 35.69s1.596-5.02 5.288-5.628m13.919.61a130.401 130.401 0 0 0 3.687 3.299c-1.83.332-3.132 1.413-3.132 1.413m11.558-3.104c-1.386 1.22-3.493 2.44-3.493 2.44c.832.747 1.47 1.912 1.47 1.912m-14.303-21.73l6.375-1.303M12.295 16.62l6.375.942m8.398-.942c.222 1.192.388 2.273.416 3.381m-11.795-2.879s.099 2.353.293 3.877"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.982 34.193s.776.692 2.19 1.718"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.77 33.901s-1.179 1.677-1.525 2.01m-5.294-1.275c.471.513 1.802 1.65 1.802 1.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.517 34.428c-.416.873-1.4 2.107-1.4 2.107"/>`),
		g.Group(children),
	)
}