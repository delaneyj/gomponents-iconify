package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeftJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 20h-13a.5.5 0 0 0 0 1h13a.5.5 0 0 0 0-1zM2.5 5h19a.5.5 0 0 0 0-1h-19a.5.5 0 0 0 0 1zm19 7h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm0-4h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1zm0 8h-19a.5.5 0 0 0 0 1h19a.5.5 0 0 0 0-1z"/>`),
		g.Group(children),
	)
}