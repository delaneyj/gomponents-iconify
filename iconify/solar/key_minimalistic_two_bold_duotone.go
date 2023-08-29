package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyMinimalisticTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15 16a7 7 0 1 0 0-14a7 7 0 0 0 0 14Zm0-5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M10.609 14.452a7.045 7.045 0 0 1-1.06-1.06l-6.58 6.578a.75.75 0 1 0 1.061 1.06l.47-.47l.97.97a.75.75 0 0 0 1.06-1.06l-.97-.97l.94-.94l.97.97a.75.75 0 0 0 1.06-1.06l-.97-.97l3.049-3.048Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}