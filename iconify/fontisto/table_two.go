package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0zM11.556 16v-2.667h6.222V16zm6.222 2.667v2.667h-6.222v-2.667zm0-10.667v2.667h-6.222V8.001zm2.667 0h6.222v2.667h-6.222zM8.889 8v2.667H2.667V8.001zm-6.222 5.334H8.89v2.667H2.667zm17.778 0h6.222v2.667h-6.222zM2.667 18.667H8.89v2.667H2.667zm17.778 2.666v-2.667h6.222v2.667z"/>`),
		g.Group(children),
	)
}