package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeedleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m206.63 94.63l-24 24C128 128 40 216 40 216s88-88 97.37-142.63l24-24a32 32 0 0 1 45.26 45.26Z" opacity=".2"/><path d="M189.66 66.34a8 8 0 0 1 0 11.32l-16 16a8 8 0 0 1-11.32-11.32l16-16a8 8 0 0 1 11.32 0ZM224 72a39.71 39.71 0 0 1-11.72 28.28l-24 24a8 8 0 0 1-4.3 2.23c-51.49 8.84-137.46 94.28-138.32 95.15a8 8 0 0 1-11.31-11.32C36 208.73 120.69 123.28 129.49 72a8 8 0 0 1 2.23-4.3l24-24A40 40 0 0 1 224 72Zm-16 0a24 24 0 0 0-41-17l-22.23 22.29c-4.41 21.15-18.9 46.19-35.49 69.43c23.24-16.59 48.28-31.08 69.43-35.49L201 89a23.85 23.85 0 0 0 7-17Z"/></g>`),
		g.Group(children),
	)
}