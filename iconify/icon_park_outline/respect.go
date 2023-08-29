package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Respect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 42H31l-2.956-13.793a2.002 2.002 0 0 1-.044-.419V15l-7.43-5.964a1.962 1.962 0 0 1 2.19-3.24L32 11l3.947 14.803c.035.131.084.258.144.38L44 42Zm-8-16l-8 2"/><path d="m19.379 25.52l.064.054a2 2 0 0 0 3-.508L24 22.463V17.77L20.32 15l-6.36.846l-4.145 8.096c-.05.096-.092.197-.125.3L4 42h11.96V29L17 24.087L17.5 22H19l-.304 1.617a2 2 0 0 0 .683 1.904ZM10 25l6 3"/></g>`),
		g.Group(children),
	)
}