package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CicsCmas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.334 30H26v-2h4v-2h-2a2.002 2.002 0 0 1-2-2v-2.334A1.668 1.668 0 0 1 27.666 20H32v2h-4v2h2a2.002 2.002 0 0 1 2 2v2.333A1.668 1.668 0 0 1 30.334 30zm-8-10h-2.667A1.667 1.667 0 0 0 18 21.666V30h2v-4h2v4h2v-8.334A1.667 1.667 0 0 0 22.334 20zM20 24v-2h2v2zm-7.5 0L11 20H9v10h2v-7l1.5 4l1.5-4v7h2V20h-2l-1.5 4zM1 22v6.5A1.473 1.473 0 0 0 2.5 30H7v-2H3v-6h4v-2H3a2.006 2.006 0 0 0-2 2zM13 8h-2v3H8v2h3v3h2v-3h3v-2h-3V8z"/><path fill="currentColor" d="M6 6h20v12h2V6a2.006 2.006 0 0 0-2-2H6a2.006 2.006 0 0 0-2 2v12h2Z"/>`),
		g.Group(children),
	)
}