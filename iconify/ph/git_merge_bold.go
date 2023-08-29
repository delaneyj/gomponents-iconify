package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 108a36.07 36.07 0 0 0-33.94 24H152a36.16 36.16 0 0 1-28.8-14.4L99.62 86.16A36 36 0 1 0 68 89.94v76.12a36 36 0 1 0 24 0V116l12 16a60.27 60.27 0 0 0 48 24h22.06A36 36 0 1 0 208 108ZM80 44a12 12 0 1 1-12 12a12 12 0 0 1 12-12Zm0 168a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm128-56a12 12 0 1 1 12-12a12 12 0 0 1-12 12Z"/>`),
		g.Group(children),
	)
}