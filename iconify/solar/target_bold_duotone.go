package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10Z" opacity=".5"/><path d="M9.25 12a.75.75 0 0 1 .75-.75h1.25V10a.75.75 0 0 1 1.5 0v1.25H14a.75.75 0 0 1 0 1.5h-1.25V14a.75.75 0 0 1-1.5 0v-1.25H10a.75.75 0 0 1-.75-.75Zm-7.222.75a10.14 10.14 0 0 1 0-1.5H5a.75.75 0 0 1 0 1.5H2.028Zm10.722 9.222a10.119 10.119 0 0 1-1.5 0V19a.75.75 0 0 1 1.5 0v2.972Zm9.222-10.722a10.119 10.119 0 0 1 0 1.5H19a.75.75 0 0 1 0-1.5h2.972ZM12.75 2.028V5a.75.75 0 0 1-1.5 0V2.028a10.14 10.14 0 0 1 1.5 0Z"/></g>`),
		g.Group(children),
	)
}