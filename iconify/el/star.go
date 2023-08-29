package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M961.359 1173.121L594.085 903.799l-374.367 259.374L362.365 730.65L0 454.756l455.436 2.008L605.848 26.879l138.827 433.765L1200 470.853L830.365 736.927l130.994 436.194z"/>`),
		g.Group(children),
	)
}