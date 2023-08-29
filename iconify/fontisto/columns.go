package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0zM2.664 8.001h10.669v13.333H2.666zM16 21.333V8h10.667v13.333z"/>`),
		g.Group(children),
	)
}