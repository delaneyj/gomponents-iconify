package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 10h-4V2h-2v8h-9V8h-2v2H8V8H6v2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2ZM4 28V12h24v16Z"/><path fill="currentColor" d="M10 26a4 4 0 1 1 4-4a4 4 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2 2 0 0 0-2-2zm-3-6h6v2H7zm10 2h9v2h-9zm0 4h9v2h-9zm0 4h9v2h-9z"/>`),
		g.Group(children),
	)
}