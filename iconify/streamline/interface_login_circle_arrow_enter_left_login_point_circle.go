package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceLoginCircleArrowEnterLeftLoginPointCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7h-8m2-2l-2 2l2 2"/><path d="M12.48 10.5a6.5 6.5 0 1 1 0-7"/></g>`),
		g.Group(children),
	)
}