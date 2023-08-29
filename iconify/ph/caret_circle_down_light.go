package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretCircleDownLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm0 192a90 90 0 1 1 90-90a90.1 90.1 0 0 1-90 90Zm44.24-110.24a6 6 0 0 1 0 8.48l-40 40a6 6 0 0 1-8.48 0l-40-40a6 6 0 0 1 8.48-8.48L128 143.51l35.76-35.75a6 6 0 0 1 8.48 0Z"/>`),
		g.Group(children),
	)
}