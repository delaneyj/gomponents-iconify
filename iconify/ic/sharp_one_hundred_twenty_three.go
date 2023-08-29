package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpOneHundredTwentyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15H5.5v-4.5H4V9h3v6zm6.5-1.5h-3v-1h3V9H9v1.5h3v1H9V15h4.5v-1.5zm6 1.5V9H15v1.5h3v1h-2v1h2v1h-3V15h4.5z"/>`),
		g.Group(children),
	)
}