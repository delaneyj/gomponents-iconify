package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 2.07A7.002 7.002 0 0 0 5 9v1h6V2.07zM5 12v3a7 7 0 1 0 14 0v-3H5zm14-2V9a7.002 7.002 0 0 0-6-6.93V10h6z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}