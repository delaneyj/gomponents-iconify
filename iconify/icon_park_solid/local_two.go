package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSLocalTwo0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44s15-12 15-25c0-8.284-6.716-15-15-15c-8.284 0-15 6.716-15 15c0 13 15 25 15 25Z"/><path fill="#000" stroke="#000" d="M24 25a6 6 0 1 0 0-12a6 6 0 0 0 0 12Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSLocalTwo0)"/>`),
		g.Group(children),
	)
}