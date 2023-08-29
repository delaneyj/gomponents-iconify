package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridDrawingForArtist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.194 6.131l.13 35.738M34.887 8.312l.11 30.434M13.395 8.345l.104 30.658M6.972 24.13h34.056M9.175 35.395h29.78M9.104 13.382h29.532m-10.215 6.845a4.281 4.281 0 1 1 .003-.01m4.663 10.942c-1.43 2.31-6.506 3.63-11.339 2.947s-7.596-3.108-6.172-5.42s6.495-3.634 11.33-2.955s7.604 3.104 6.187 5.416"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}