package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 208H48V48a160 160 0 0 1 160 160Z" opacity=".2"/><path d="M98.91 157.09A71.53 71.53 0 0 1 120 208a8 8 0 0 1-16 0a56 56 0 0 0-56-56a8 8 0 0 1 0-16a71.53 71.53 0 0 1 50.91 21.09ZM48 88a8 8 0 0 0 0 16a104 104 0 0 1 104 104a8 8 0 0 0 16 0A120 120 0 0 0 48 88Zm118.79 1.21A166.89 166.89 0 0 0 48 40a8 8 0 0 0 0 16a151 151 0 0 1 107.48 44.52A151 151 0 0 1 200 208a8 8 0 0 0 16 0a166.9 166.9 0 0 0-49.21-118.79ZM52 192a12 12 0 1 0 12 12a12 12 0 0 0-12-12Z"/></g>`),
		g.Group(children),
	)
}