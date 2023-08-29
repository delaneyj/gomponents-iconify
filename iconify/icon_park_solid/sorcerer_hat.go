package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SorcererHat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSorcererHat0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 40c11.046 0 20-1.79 20-4c0-1.439-3.299-2.7-9-3.405L27 12L15 8l3 6l-5 18.595C7.299 33.3 4 34.56 4 36c0 2.21 8.954 4 20 4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSorcererHat0)"/>`),
		g.Group(children),
	)
}