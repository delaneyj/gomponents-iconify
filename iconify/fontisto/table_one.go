package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0zM13.334 8.001v5.333H2.667V8.001zm2.666 0h10.667v5.333H16zM2.667 16h10.667v5.333H2.667zM16 21.333V16h10.667v5.333z"/>`),
		g.Group(children),
	)
}