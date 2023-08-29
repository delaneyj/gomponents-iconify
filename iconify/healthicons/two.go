package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Two(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M26 14a4 4 0 0 1 2.61 7.034l-14 13.528A2 2 0 0 0 16 38h16a2 2 0 1 0 0-4H20.948l10.416-10.065A8 8 0 0 0 26 10h-4a8.003 8.003 0 0 0-7.544 5.334a2 2 0 0 0 3.771 1.332A4.002 4.002 0 0 1 22 14h4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}