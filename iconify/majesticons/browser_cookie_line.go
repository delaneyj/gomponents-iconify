package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserCookieLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8V6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v2M3 8v10a2 2 0 0 0 2 2h4M3 8h18m0 0v2m-6 7h.001M18 19h.001M22 17a5 5 0 1 1-5-5v2h3v3h2z"/>`),
		g.Group(children),
	)
}