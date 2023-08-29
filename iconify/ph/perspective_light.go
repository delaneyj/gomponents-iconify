package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PerspectiveLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 122h-18V48a14 14 0 0 0-16.5-13.77l-160 29.09A14 14 0 0 0 34 77.09V122H16a6 6 0 0 0 0 12h18v44.91a14 14 0 0 0 11.5 13.77l160 29.09a14.2 14.2 0 0 0 2.5.23a14 14 0 0 0 14-14v-74h18a6 6 0 0 0 0-12ZM46 77.09a2 2 0 0 1 1.64-2l160-29.1h.37a2 2 0 0 1 2 2v74H46ZM210 208a2 2 0 0 1-.72 1.53a2 2 0 0 1-1.64.44l-160-29.1a2 2 0 0 1-1.64-2V134h164Z"/>`),
		g.Group(children),
	)
}