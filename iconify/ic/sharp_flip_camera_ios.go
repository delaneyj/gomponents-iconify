package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpFlipCameraIos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.83 5L15 3H9L7.17 5H2v16h20V5h-5.17zM12 18c-2.76 0-5-2.24-5-5H5l2.5-2.5L10 13H8c0 2.21 1.79 4 4 4c.58 0 1.13-.13 1.62-.35l.74.74c-.71.37-1.5.61-2.36.61zm4.5-2.5L14 13h2c0-2.21-1.79-4-4-4c-.58 0-1.13.13-1.62.35l-.74-.73C10.35 8.24 11.14 8 12 8c2.76 0 5 2.24 5 5h2l-2.5 2.5z"/>`),
		g.Group(children),
	)
}