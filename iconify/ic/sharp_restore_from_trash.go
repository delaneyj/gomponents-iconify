package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpRestoreFromTrash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21h12V7H6v14zm6-11l4 4h-2v4h-4v-4H8l4-4zm3.5-6l-1-1h-5l-1 1H5v2h14V4z"/>`),
		g.Group(children),
	)
}