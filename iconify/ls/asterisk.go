package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m311 59l37 62l-158 58l158 60l-37 62l-129-108l28 167h-73l30-167L36 301L0 239l160-60L0 121l36-62l131 108L137 0h73l-28 167z"/>`),
		g.Group(children),
	)
}