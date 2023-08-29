package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 12a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1zm3-4h15a.5.5 0 0 0 0-1h-15a.5.5 0 0 0 0 1zm-3 9a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1zm18-5h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1zm-18-5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1zm18 10h-15a.5.5 0 0 0 0 1h15a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}