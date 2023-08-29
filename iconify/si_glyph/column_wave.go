package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnWave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 2.057c0-.584.447-1.057.999-1.057S16 1.473 16 2.057v12.719c0 .584-.449 1.057-1.001 1.057c-.552 0-.999-.473-.999-1.057V2.057zm-4 3c0-.584.449-1.057.999-1.057C11.552 4 12 4.473 12 5.057v9.719c0 .584-.448 1.057-1.001 1.057c-.55 0-.999-.473-.999-1.057V5.057zm-3.998 3C6.002 7.473 6.449 7 7 7c.552 0 1 .473 1 1.057v6.719c0 .584-.448 1.057-1 1.057c-.551 0-.998-.473-.998-1.057V8.057zm-3.959 3c0-.584.438-1.057.978-1.057S4 10.473 4 11.057v3.719c0 .584-.439 1.057-.979 1.057s-.978-.473-.978-1.057v-3.719z"/>`),
		g.Group(children),
	)
}