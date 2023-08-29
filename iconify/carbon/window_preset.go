package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowPreset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M15 6h2v2h-2z" fill="currentColor"/><path d="M24 15h2v2h-2z" fill="currentColor"/><path d="M15 24h2v2h-2z" fill="currentColor"/><path d="M6 15h2v2H6z" fill="currentColor"/><path d="M8.172 9.586l1.414-1.414L11 9.586L9.586 11z" fill="currentColor"/><path d="M21 9.586l1.414-1.414l1.414 1.414L22.414 11z" fill="currentColor"/><path d="M21 22.414L22.414 21l1.414 1.414l-1.414 1.414z" fill="currentColor"/><path d="M8.172 22.414L9.586 21L11 22.414l-1.414 1.414z" fill="currentColor"/><path d="M16 22a6 6 0 1 0-6-6a6 6 0 0 0 6 6zm0-10v8a4 4 0 0 1 0-8z" fill="currentColor"/><path d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2zM4 4v24h24V4z" fill="currentColor"/>`),
		g.Group(children),
	)
}