package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soccer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSoccer0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M17.817 4.98C7.31 8.39 1.57 19.677 4.98 30.176c3.41 10.498 14.698 16.247 25.196 12.838c10.508-3.41 16.247-14.698 12.838-25.196C39.603 7.309 28.315 1.57 17.817 4.98Z"/><path fill="#fff" d="m34 21l-10-8l-10 8l4 12h12l4-12Z"/><path d="m34 21l9-3m-7 22l-6-7m-12 0l-6 7m2-19l-9-3m19-5V4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSoccer0)"/>`),
		g.Group(children),
	)
}