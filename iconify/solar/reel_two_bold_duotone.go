package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelTwoBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 12c0 5.523 4.477 10 10 10h9.25a.75.75 0 0 0 0-1.5h-3.98A9.993 9.993 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12Z" opacity=".5"/><path fill-rule="evenodd" d="M12 15.75a3.75 3.75 0 1 1 0-7.5a3.75 3.75 0 0 1 0 7.5Z" clip-rule="evenodd"/><path d="M5.5 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM12 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm1 14a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm5.5-5.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/></g>`),
		g.Group(children),
	)
}