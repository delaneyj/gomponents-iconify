package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mosaic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" d="M44 36h-8v8h8v-8Zm-16 0h-8v8h8v-8Zm-16 0H4v8h8v-8Zm32-16h-8v8h8v-8Zm-16 0h-8v8h8v-8Zm-16 0H4v8h8v-8ZM44 4h-8v8h8V4ZM28 4h-8v8h8V4ZM12 4H4v8h8V4Zm8 8h-8v8h8v-8Zm0 16h-8v8h8v-8Zm16-16h-8v8h8v-8Zm0 16h-8v8h8v-8Z"/>`),
		g.Group(children),
	)
}