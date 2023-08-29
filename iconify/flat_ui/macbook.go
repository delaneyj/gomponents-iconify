package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Macbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#2C3E50" fill-rule="evenodd" d="M10 62V5a5 5 0 0 1 5-5h70a5 5 0 0 1 5 5v57H10z" clip-rule="evenodd"/><path fill="#35495E" fill-rule="evenodd" d="M14 4h72v54H14V4z" clip-rule="evenodd"/><path fill="#BDC3C7" fill-rule="evenodd" d="M0 62h100v3c0 2.209-2.791 3-5 3H5c-2.209 0-5-.791-5-3v-3z" clip-rule="evenodd"/><path fill="#95A5A5" fill-rule="evenodd" d="M59.95 62a2.5 2.5 0 0 1-2.45 2h-15a2.502 2.502 0 0 1-2.45-2h19.9z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}