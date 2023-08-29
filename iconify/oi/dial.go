package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M4 1C1.8 1 0 2.8 0 5h1c0-1.66 1.34-3 3-3s3 1.34 3 3h1c0-2.2-1.8-4-4-4zm-.59 2.09A2 2 0 0 0 4 7c1.11 0 2-.89 2-2c0-.9-.59-1.65-1.41-1.91L4 3.97l-.59-.88z"/>`),
		g.Group(children),
	)
}