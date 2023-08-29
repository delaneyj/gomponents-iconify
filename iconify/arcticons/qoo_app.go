package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QooApp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.973 16.436a14.054 14.054 0 0 1 2.411 7.894c0 14.16-14.12 15.027-14.12 19.17c0-4.143-14.12-5.01-14.12-19.17c0-7.798 6.322-14.12 14.12-14.12c.775 0 1.535.063 2.275.183"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.264 26.388c1.227 1.493 2.916 2.5 5.03 2.5c8.17 0 6.559-6.502 6.559-6.502s.307 1.944-2.416 1.944c-1.374 0-3.469-2.614-5.274-2.614c-1.764 0-3.055.556-3.899 2.033c-.843-1.477-2.134-2.033-3.898-2.033c-1.805 0-3.9 2.614-5.274 2.614c-2.723 0-2.416-1.944-2.416-1.944s-1.61 6.502 6.558 6.502c2.115 0 3.804-1.007 5.03-2.5Zm.25-17.303l14.242 9.205m-11.002-7.11S28.96 6.17 30.01 4.5c3.247 1.16 7.62 4.816 8.847 6.861c-1.714 2.046-4.34 4.835-4.34 4.835"/>`),
		g.Group(children),
	)
}