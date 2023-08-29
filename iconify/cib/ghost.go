package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12.807 25.599H.01V32h12.797zm19.183 0H19.204V32H31.99zm.005-12.797H0v6.396h31.995zM32 0h-6.401v6.401H32zM19.203 0H.01v6.401h19.193z"/>`),
		g.Group(children),
	)
}