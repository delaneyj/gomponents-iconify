package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRss0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M8 44V6a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2v38l-16-8.273L8 44Z"/><path stroke="#000" stroke-linecap="round" d="M23.95 13.95v12m-6-6h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRss0)"/>`),
		g.Group(children),
	)
}