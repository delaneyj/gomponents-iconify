package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForwardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M248 128a15.76 15.76 0 0 1-7.33 13.34l-88.19 56.16A15.91 15.91 0 0 1 128 184.16v-37.3L48.48 197.5A15.91 15.91 0 0 1 24 184.16V71.84A15.91 15.91 0 0 1 48.48 58.5L128 109.14v-37.3a15.91 15.91 0 0 1 24.48-13.34l88.19 56.16A15.76 15.76 0 0 1 248 128Z"/>`),
		g.Group(children),
	)
}