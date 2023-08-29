package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileImage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3V1Zm2 2v18h14V9h-6V3H5Zm10 .414V7h3.586L15 3.414ZM8 11h2.004v2.004H8V11Zm4.728 1.938L16.79 17l-1.414 1.414l-2.648-2.648l-1.612 1.613l-.652-.784l-.7.677l-1.142 1.142L7.208 17l1.154-1.154l2.262-2.188l.628.756l1.476-1.476Z"/>`),
		g.Group(children),
	)
}