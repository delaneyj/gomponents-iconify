package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleSafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPeopleSafe0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path d="M6 9.256L24.009 4L42 9.256v10.778C42 31.362 34.75 40.419 24.003 44C13.253 40.42 6 31.36 6 20.029V9.256Z"/><circle cx="24" cy="18" r="5" fill="#fff" stroke-linecap="round"/><path stroke-linecap="round" d="M32 31a8 8 0 1 0-16 0"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPeopleSafe0)"/>`),
		g.Group(children),
	)
}