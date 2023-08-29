package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="m27.755 31.724l15.267 15.264a3 3 0 1 0 4.242-4.243l-23.069-23.07a6 6 0 1 0-8.485 8.485l25.198 25.305c3.906 3.905 10.237 3.905 14.142 0c3.906-3.905 3.906-10.237 0-14.142L42.547 26.82"/>`),
		g.Group(children),
	)
}