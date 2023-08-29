package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XxOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#fff" fill-rule="evenodd" stroke="#adb5bd" d="M.5.5h511v511H.5z"/><path fill="none" stroke="#adb5bd" d="m.5.5l511 511m0-511l-511 511"/>`),
		g.Group(children),
	)
}