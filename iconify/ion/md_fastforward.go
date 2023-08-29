package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdFastforward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M480 256L262.4 110v292L480 256zM32 110v292l217.6-146L32 110z" fill="currentColor"/>`),
		g.Group(children),
	)
}