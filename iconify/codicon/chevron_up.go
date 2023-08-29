package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.024 5.928l-4.357 4.357l-.62-.618L7.716 5h.618L13 9.667l-.619.618l-4.357-4.357z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}