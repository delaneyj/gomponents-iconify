package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 24h9v2h-9zm1 4h6v2h-6z"/><path fill="currentColor" d="m8.783 18.975l1.418-1.418A6.92 6.92 0 0 1 8 12a7.99 7.99 0 0 1 13.527-5.77l1.414-1.413A9.988 9.988 0 0 0 6 12a8.941 8.941 0 0 0 2.783 6.975zM30 3.415L28.586 2L2 28.586L3.414 30L23.661 9.753A7.707 7.707 0 0 1 24 12a7.203 7.203 0 0 1-2.822 6.143C20.107 19.135 19 20.16 19 22h2c0-.92.526-1.45 1.535-2.386A9.193 9.193 0 0 0 26 12a9.874 9.874 0 0 0-.776-3.81z"/>`),
		g.Group(children),
	)
}