package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileShareTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<g id="clarityFileShare2Line0" fill="currentColor"><path d="M25 4H7.83A1.89 1.89 0 0 0 6 5.91v24.18A1.89 1.89 0 0 0 7.83 32h20.34A1.87 1.87 0 0 0 30 30.09V9Zm-1 1.78L28.2 10H24ZM8 30V6h14v6h6v18Z"/><path d="M22 21.81a2.11 2.11 0 0 0-1.44.62l-5.72-2.66v-.44l5.66-2.65a2.08 2.08 0 1 0 .06-2.94a2.14 2.14 0 0 0-.64 1.48v.23l-5.64 2.66a2.08 2.08 0 1 0-.08 2.95l.08-.08l5.67 2.66v.3a2.09 2.09 0 1 0 2.05-2.1Z"/></g>`),
		g.Group(children),
	)
}