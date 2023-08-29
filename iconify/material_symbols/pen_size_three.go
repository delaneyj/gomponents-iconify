package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenSizeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.45 18.55q-.425-.425-.425-1.05t.425-1.05l11-11q.425-.45 1.05-.438t1.05.438q.425.425.438 1.05t-.438 1.05l-11 11q-.425.425-1.05.438t-1.05-.438Z"/>`),
		g.Group(children),
	)
}