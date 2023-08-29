package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningFourOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.75 4a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0ZM14 5.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm-2 9.5a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5ZM10.75 12a1.25 1.25 0 1 0 2.5 0a1.25 1.25 0 0 0-2.5 0ZM10 22.75a2.75 2.75 0 1 1 0-5.5a2.75 2.75 0 0 1 0 5.5ZM8.75 20a1.25 1.25 0 1 0 2.5 0a1.25 1.25 0 0 0-2.5 0Z" clip-rule="evenodd"/><path d="M15.25 12a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 0 1.5h-3a.75.75 0 0 1-.75-.75ZM14 19.25a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 0-1.5h-5ZM10.75 4a.75.75 0 0 0-.75-.75H5a.75.75 0 0 0 0 1.5h5a.75.75 0 0 0 .75-.75ZM5 11.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 0-1.5H5ZM4.25 20a.75.75 0 0 1 .75-.75h1a.75.75 0 0 1 0 1.5H5a.75.75 0 0 1-.75-.75ZM19 3.25a.75.75 0 0 1 0 1.5h-1a.75.75 0 0 1 0-1.5h1Z"/></g>`),
		g.Group(children),
	)
}