package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 0v1200h-200V650l-500 500V650L0 1150V50l500 500V50l500 500V0h200z"/>`),
		g.Group(children),
	)
}