package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6H7v16h8v-2h2v-2h-2v-2h2v2h2v-2h2V6zM9 20V8h10v6h-6v6H9zm-6-2h2V4h12V2H3v16z"/>`),
		g.Group(children),
	)
}