package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 16a7 7 0 1 0-5.452-2.609L3.97 18.97l-1 1a.75.75 0 1 0 1.06 1.061l.47-.47l.97.97a.75.75 0 0 0 1.06-1.06l-.97-.97l.94-.94l.97.97a.75.75 0 0 0 1.06-1.06l-.97-.97l3.049-3.048A6.971 6.971 0 0 0 15 16Zm0-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}