package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaiduNetdisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="13.81" cy="29.904" r="9.307" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.441" cy="17.321" r="8.532" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.14 35.639l11.161-12.117c3.288-3.568 12.447-1.495 12.447 5.863c0 9.826-8.692 13.045-15.019 4.747m-4.597-20.767c-2.054 1.236-3.8 2.713-2.62 6.675"/>`),
		g.Group(children),
	)
}