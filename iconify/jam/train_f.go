package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v5.038l12-.023V4a2 2 0 0 0-2-2H4zm7.667 15L11 16H5l-.667 1H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v9a4 4 0 0 1-4 4h-.333zM4 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm0-11a1 1 0 0 1 1 1v3a1 1 0 0 1-2 0V4a1 1 0 0 1 1-1zM4.314 17.029l7.371-.001l.279.417A1 1 0 0 1 11.13 19H4.87a1 1 0 0 1-.833-1.555l.278-.416z"/>`),
		g.Group(children),
	)
}