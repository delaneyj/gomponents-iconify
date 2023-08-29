package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.016 0v1.031h-1.062V0h-.895v1.031h-1.09V0h-.953v1.031H9.954V0h-.922v1.031H7.941V0h-.925v1.031H5.985V0h-.942v1.031H3.959V0H3v16h12.954V0h-.938zM5 6.958h9v1H5v-1zm9 5H5v-1h9v1zM14 10H5V9h9v1z"/>`),
		g.Group(children),
	)
}