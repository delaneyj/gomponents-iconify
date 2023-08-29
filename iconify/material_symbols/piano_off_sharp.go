package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PianoOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L18.15 21H3V5.85L1.375 4.225L2.8 2.8l18.4 18.4l-1.425 1.425ZM21 18.15l-2-2V5h-2v9.15l-4-4V5h-2v3.15L5.85 3H21v15.15ZM5 19h3.25v-4.5H7V9.85l-2-2V19Zm4.75 0h4.5v-1.9l-3.3-3.3v.7h-1.2V19Zm6 0h.4l-.4-.4v.4Z"/>`),
		g.Group(children),
	)
}