package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightToLeftTextDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 17H5.91l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 19H20.5a1 1 0 0 0 0-2Zm-9-7v4a1 1 0 0 0 2 0V4h2v10a1 1 0 0 0 2 0V4h1a1 1 0 0 0 0-2h-7a4 4 0 0 0 0 8Zm0-6v4a2 2 0 0 1 0-4Z"/>`),
		g.Group(children),
	)
}