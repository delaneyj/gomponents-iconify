package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadRetroF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 5V4a1 1 0 1 0-2 0v1H4a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0V7h1a1 1 0 1 0 0-2H7zM6 0h12a6 6 0 1 1-4.472 10h-3.056A6 6 0 1 1 6 0zm12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-2 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm4 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-2 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}