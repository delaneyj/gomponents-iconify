package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M18.25 7a3.75 3.75 0 1 1-7.5 0a3.75 3.75 0 0 1 7.5 0ZM14.5 9.25a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5ZM5.75 17a3.75 3.75 0 1 0 7.5 0a3.75 3.75 0 0 0-7.5 0Zm3.75 2.25a2.25 2.25 0 1 1 0-4.5a2.25 2.25 0 0 1 0 4.5Z" clip-rule="evenodd"/><path d="M14.25 16.959a.75.75 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5h-7a.75.75 0 0 1-.75-.75ZM9 6.209a.75.75 0 0 1 0 1.5H2a.75.75 0 0 1 0-1.5h7Zm-7.75 10.75a.75.75 0 0 1 .75-.75h2a.75.75 0 0 1 0 1.5H2a.75.75 0 0 1-.75-.75ZM22 6.209a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h2Z"/></g>`),
		g.Group(children),
	)
}