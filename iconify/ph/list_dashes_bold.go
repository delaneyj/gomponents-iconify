package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListDashesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 64a12 12 0 0 1 12-12h120a12 12 0 0 1 0 24H96a12 12 0 0 1-12-12Zm132 52H96a12 12 0 0 0 0 24h120a12 12 0 0 0 0-24Zm0 64H96a12 12 0 0 0 0 24h120a12 12 0 0 0 0-24ZM56 52H40a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Zm0 64H40a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Zm0 64H40a12 12 0 0 0 0 24h16a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}