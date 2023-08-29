package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterCloudFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10a1 1 0 0 1 1 1v3H6V3a1 1 0 0 1 1-1Zm15 7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h2v-5h7.194c.243-.891.715-1.688 1.417-2.39C14.685 11.538 15.98 11 17.5 11c1.519 0 2.815.537 3.89 1.61c.227.229.43.466.61.714V9ZM8 10v2H5v-2h3Zm13 6.5a3.5 3.5 0 1 0-7 0l.003.102a2.751 2.751 0 0 0 .58 5.393l.167.005h5.5l.168-.005a2.75 2.75 0 0 0 .58-5.392L21 16.5ZM7 17h3.562A4.617 4.617 0 0 0 10 19.25c0 1.032.29 1.949.871 2.75H7v-5Z"/>`),
		g.Group(children),
	)
}