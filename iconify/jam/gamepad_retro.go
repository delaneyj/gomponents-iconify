package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadRetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5h1a1 1 0 1 1 0 2H7v1a1 1 0 1 1-2 0V7H4a1 1 0 1 1 0-2h1V4a1 1 0 1 1 2 0v1zM6 0h12a6 6 0 1 1-4.472 10h-3.056A6 6 0 1 1 6 0zm0 2a4 4 0 1 0 2.982 6.666L9.578 8h4.844l.596.666A4 4 0 1 0 18 2H6zm12 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm-2 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm4 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm-2 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}