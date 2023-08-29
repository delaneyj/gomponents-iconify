package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 1.75a.75.75 0 0 0-1.5 0V6.5a.5.5 0 0 1-1 0V2.75a.75.75 0 0 0-1.5 0V7.5a.5.5 0 0 1-1 0V4.75A.75.75 0 0 0 2 4.748V9.5c0 .813.344 1.71.743 2.492c.407.797.906 1.54 1.283 2.059c.45.622 1.171.949 1.91.949h2.328c.952 0 1.797-.54 2.255-1.34a35.659 35.659 0 0 1 2.233-3.38a40.8 40.8 0 0 1 1.112-1.435l.015-.02l.004-.004h.001a.5.5 0 0 0-.03-.675a1.578 1.578 0 0 0-1.187-.482c-.4.01-.778.159-1.096.336a4.663 4.663 0 0 0-.57.379V2.75a.75.75 0 0 0-1.5 0V6.5a.5.5 0 0 1-1 0V1.75Z"/>`),
		g.Group(children),
	)
}