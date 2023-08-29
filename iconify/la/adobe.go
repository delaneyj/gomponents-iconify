package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 6v20l8.9-20H4zm15.1 0L28 26V6h-8.9zM16 13.4L12.1 22h4.097l1.6 4H21.6L16 13.4z"/>`),
		g.Group(children),
	)
}