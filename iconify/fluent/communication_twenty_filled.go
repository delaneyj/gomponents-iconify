package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommunicationTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.5 11a6.5 6.5 0 1 1 11.101 4.591a.75.75 0 1 0 1.062 1.06a8 8 0 1 0-11.326 0A.75.75 0 0 0 5.4 15.59A6.477 6.477 0 0 1 3.5 11Zm3 0a3.5 3.5 0 1 1 5.98 2.47a.75.75 0 1 0 1.062 1.06a5 5 0 1 0-7.083 0a.75.75 0 0 0 1.062-1.06A3.487 3.487 0 0 1 6.5 11ZM10 9a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}