package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextIndentLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a6 6 0 0 1-6 6H112a6 6 0 0 1 0-12h104a6 6 0 0 1 6 6ZM112 70h104a6 6 0 0 0 0-12H112a6 6 0 0 0 0 12Zm104 116H40a6 6 0 0 0 0 12h176a6 6 0 0 0 0-12ZM35.76 140.24a6 6 0 0 0 8.48 0l40-40a6 6 0 0 0 0-8.48l-40-40a6 6 0 0 0-8.48 8.48L71.51 96l-35.75 35.76a6 6 0 0 0 0 8.48Z"/>`),
		g.Group(children),
	)
}