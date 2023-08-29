package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0L0 1200h1200L600 0zm-46.143 416.089h92.284v158.644l-22.559 223.096h-47.168l-22.56-223.096V416.089h.003zm0 469.336h92.284v99.391h-92.284v-99.391z"/>`),
		g.Group(children),
	)
}