package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotesDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H3v20h10v-2H5V4h14v10h2V2h-2zm-2 4H7v2h10V6zM7 10h10v2H7v-2zm6 4H7v2h6v-2zm6 4h-2v-2h-2v2h2v2h-2v2h2v-2h2v2h2v-2h-2v-2zm0 0h2v-2h-2v2z"/>`),
		g.Group(children),
	)
}