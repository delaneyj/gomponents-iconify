package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAdobeIllustrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm4.756 2h2.488l3.6 12h-2.091l-1.183-3.945H8.428L7.244 18H5.153L8.756 6Zm.272 6.055h1.943L10 8.815l-.972 3.24ZM15.996 9H18v2.004h-2.004V9ZM18 12v6h-2v-6h2Z"/>`),
		g.Group(children),
	)
}