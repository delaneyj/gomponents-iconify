package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextOutdentLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 128a6 6 0 0 1-6 6H112a6 6 0 0 1 0-12h104a6 6 0 0 1 6 6ZM112 70h104a6 6 0 0 0 0-12H112a6 6 0 0 0 0 12Zm104 116H40a6 6 0 0 0 0 12h176a6 6 0 0 0 0-12ZM72 142a6 6 0 0 0 4.24-10.24L40.49 96l35.75-35.76a6 6 0 0 0-8.48-8.48l-40 40a6 6 0 0 0 0 8.48l40 40A6 6 0 0 0 72 142Z"/>`),
		g.Group(children),
	)
}