package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12.926 3.66a2.254 2.254 0 0 1 3.817 1.965l-.563 3.378A5 5 0 0 1 21 14v2a5 5 0 0 1-5 5H8V8.586l4.926-4.926zM6 9a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3V9z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}