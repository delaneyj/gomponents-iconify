package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EjectLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.737 13h8.526L12 6.606L7.737 13Zm4.679-9.376l7.066 10.599a.5.5 0 0 1-.416.777H4.934a.5.5 0 0 1-.416-.777l7.066-10.599a.5.5 0 0 1 .832 0ZM5 17h14a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2Z"/>`),
		g.Group(children),
	)
}