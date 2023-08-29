package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Very(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.133 5.5l14.182 37l11.692-.088l9.86-25.786l-11.083-5.234l-5.02 14.879L19.4 5.5ZM25.77 26.27L20.31 42.5"/>`),
		g.Group(children),
	)
}