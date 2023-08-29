package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AliensLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a8.5 8.5 0 0 1 8.5 8.5c0 6.5-5.5 12-8.5 12s-8.5-5.5-8.5-12A8.5 8.5 0 0 1 12 2Zm0 2a6.5 6.5 0 0 0-6.5 6.5c0 4.794 4.165 10 6.5 10s6.5-5.206 6.5-10A6.5 6.5 0 0 0 12 4Zm5.5 7c.16 0 .319.008.475.025a4.5 4.5 0 0 1-4.95 4.95A4.5 4.5 0 0 1 17.5 11Zm-11 0a4.5 4.5 0 0 1 4.475 4.975a4.5 4.5 0 0 1-4.95-4.95C6.18 11.008 6.34 11 6.5 11Z"/>`),
		g.Group(children),
	)
}