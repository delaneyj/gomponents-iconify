package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuDotsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.25 12a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0ZM5 10.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5ZM9.25 12a2.75 2.75 0 1 1 5.5 0a2.75 2.75 0 0 1-5.5 0ZM12 10.75a1.25 1.25 0 1 0 0 2.5a1.25 1.25 0 0 0 0-2.5Zm7-1.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5ZM17.75 12a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}