package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TentLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m253.48 197.56l-64-144A6 6 0 0 0 184 50H72a6 6 0 0 0-5.45 3.51v.14L2.52 197.56A6 6 0 0 0 8 206h240a6 6 0 0 0 5.48-8.44ZM66 84.27V194H17.23ZM78 194V84.27L126.77 194Zm61.9 0L81.23 62h98.87l58.67 132Z"/>`),
		g.Group(children),
	)
}