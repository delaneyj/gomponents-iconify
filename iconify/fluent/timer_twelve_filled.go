package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 0a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1h-4ZM1 6.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0Zm4 1a.5.5 0 0 0 1 0v-3a.5.5 0 0 0-1 0v3Zm5.854-3.647a.5.5 0 0 1-.707.001l-1.002-.998a.5.5 0 1 1 .706-.708l1.002.998a.5.5 0 0 1 .001.707Z"/>`),
		g.Group(children),
	)
}