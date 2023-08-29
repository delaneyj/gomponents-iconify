package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideImageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18.15l-2-2V5H7.85l-2-2H19q.825 0 1.413.588T21 5v13.15ZM5 21q-.825 0-1.413-.588T3 19V5.8l-.9-.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9H5Zm9.175-4H7q-.3 0-.45-.275t.05-.525l2-2.675q.15-.2.4-.2t.4.2L11.25 16l.825-1.1L5 7.825V19h11.175l-2-2Zm-.75-6.425ZM10.6 13.4Z"/>`),
		g.Group(children),
	)
}