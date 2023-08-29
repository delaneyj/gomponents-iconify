package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.555 14a.5.5 0 0 0 .816.387l2.445-2a.5.5 0 0 0 0-.773l-2.445-2a.5.5 0 0 0-.633.773L20.71 12l-1.972 1.613a.5.5 0 0 0-.183.386zM2.5 6.5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm0 4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm0 4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm19 3h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}