package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockCountdownBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236 137A108.13 108.13 0 1 1 119 20a12 12 0 0 1 2 24a84.12 84.12 0 1 0 91 91a12 12 0 1 1 24 2ZM116 76v52a12 12 0 0 0 12 12h52a12 12 0 0 0 0-24h-40V76a12 12 0 0 0-24 0Zm92 20a16 16 0 1 0-16-16a16 16 0 0 0 16 16Zm-32-32a16 16 0 1 0-16-16a16 16 0 0 0 16 16Z"/>`),
		g.Group(children),
	)
}