package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.5 0h-9a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2Zm-9 18V2h9v16h-9Z" clip-rule="evenodd"/><path d="M10 17.25a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/></g>`),
		g.Group(children),
	)
}