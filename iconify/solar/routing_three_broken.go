package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoutingThreeBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="5" cy="5" r="3" stroke="currentColor" stroke-width="1.5"/><circle cx="19" cy="19" r="3" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="M11 4.25a.75.75 0 0 0 0 1.5v-1.5ZM13 19l.53.53a.75.75 0 0 0 0-1.06L13 19Zm4.206-10.313l.402.633l-.402-.633ZM6.794 15.313l.403.632l-.403-.632Zm5.236 1.657a.75.75 0 0 0-1.06 1.06l1.06-1.06Zm-1.06 3a.75.75 0 0 0 1.06 1.06l-1.06-1.06Zm-.567-6.064a.75.75 0 0 0-.806-1.266l.806 1.265Zm2.797-3.559a.75.75 0 0 0 .806 1.266l-.806-1.266Zm2.932-6.097H11v1.5h5.132v-1.5ZM13 18.25H7.868v1.5H13v-1.5Zm.53.22l-1.5-1.5l-1.06 1.06l1.5 1.5l1.06-1.06Zm-1.06 0l-1.5 1.5l1.06 1.06l1.5-1.5l-1.06-1.06Zm-4.602-.22c-1.25 0-1.726-1.633-.671-2.305l-.805-1.265c-2.321 1.477-1.275 5.07 1.476 5.07v-1.5Zm8.264-12.5c1.25 0 1.726 1.633.671 2.305l.805 1.265c2.321-1.477 1.275-5.07-1.476-5.07v1.5Zm-6.535 6.89l-3.205 2.04l.805 1.265l3.206-2.04l-.806-1.265Zm7.206-4.585L13.2 10.347l.806 1.266l3.602-2.293l-.805-1.265Z"/></g>`),
		g.Group(children),
	)
}