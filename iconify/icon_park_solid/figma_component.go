package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaComponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFigmaComponent0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m17 12l7-7l7 7l-7 7l-7-7Zm0 24l7-7l7 7l-7 7l-7-7Zm12-12l7-7l7 7l-7 7l-7-7ZM5 24l7-7l7 7l-7 7l-7-7Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFigmaComponent0)"/>`),
		g.Group(children),
	)
}