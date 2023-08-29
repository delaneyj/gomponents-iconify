package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M13.5 12a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0Zm-13 0A11.56 11.56 0 0 1 6.249 2l3.75 6.535A3.998 3.998 0 0 0 8 12H.5Zm9.5 3.465L6.25 21.96A11.447 11.447 0 0 0 12 23.5a11.44 11.44 0 0 0 5.75-1.54L14 15.465A3.982 3.982 0 0 1 12 16a3.98 3.98 0 0 1-2-.535ZM16 12h7.5a11.56 11.56 0 0 0-5.749-10l-3.75 6.535A4.002 4.002 0 0 1 16 12Z"/>`),
		g.Group(children),
	)
}