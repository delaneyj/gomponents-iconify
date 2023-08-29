package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1v2.5h-2V1h2Zm8.414 2.5L15 9.914L13.586 8.5L20 2.086L21.414 3.5ZM4 2.086l11.968 11.968a4.001 4.001 0 0 1 4.86 6.274a4 4 0 0 1-6.274-4.86L12 12.914l-2.554 2.554a4.002 4.002 0 0 1-6.274 4.86a4 4 0 0 1 4.86-6.274l2.554-2.554l-8-8L4 2.086ZM13 5v2.5h-2V5h2ZM7.414 16.086a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828Zm12 0a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828Z"/>`),
		g.Group(children),
	)
}