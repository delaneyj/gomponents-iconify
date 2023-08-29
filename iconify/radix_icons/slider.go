package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.3 7.5a1.8 1.8 0 1 1-3.6 0a1.8 1.8 0 0 1 3.6 0Zm.905.5a2.751 2.751 0 0 1-5.41 0H.5a.5.5 0 0 1 0-1h5.295a2.751 2.751 0 0 1 5.41 0H14.5a.5.5 0 0 1 0 1h-3.295Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}