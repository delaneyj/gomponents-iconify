package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 11h2v10H7zm4 0h2v10h-2zm4 0h2v10h-2zm4 0h2v10h-2zm4 0h2v10h-2z"/><path fill="currentColor" d="M28 26H4a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2ZM4 8v16h24V8Z"/>`),
		g.Group(children),
	)
}