package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CounterclockwiseArrowsButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M27.3 20.1c1.5-.7 3.2-1 4.9-1c6.6 0 11.9 5.1 11.9 12H52c0-10.8-8.9-19.8-19.8-19.8c-3.8 0-7.6 1.1-10.8 3.2L18 11v13.1h13.5l-4.2-4m9.2 22.5c-1.5.7-3.1 1-4.7 1c-6.6 0-12-5.3-12-12.6H12c0 11.4 8.9 20.8 19.8 20.8c3.7 0 7.2-1.1 10.3-3.1L46 53V38.2H32.6l3.9 4.4"/>`),
		g.Group(children),
	)
}