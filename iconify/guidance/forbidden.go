package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forbidden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M.5 12C.5 5.649 5.649.5 12 .5S23.5 5.649 23.5 12S18.351 23.5 12 23.5S.5 18.351.5 12Z"/><path d="M4.5 14h15v-.25l-.068-.272a6.093 6.093 0 0 1 0-2.956l.068-.272V10h-15v.25l.068.272a6.092 6.092 0 0 1 0 2.956l-.068.272V14Z"/></g>`),
		g.Group(children),
	)
}