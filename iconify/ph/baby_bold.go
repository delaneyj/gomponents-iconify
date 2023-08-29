package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M92 144a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm72-32a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm-14.4 49.85a41 41 0 0 1-43.2 0a12 12 0 1 0-12.8 20.3a65 65 0 0 0 68.8 0a12 12 0 1 0-12.8-20.3ZM236 128A108 108 0 1 1 128 20a108.12 108.12 0 0 1 108 108Zm-24 0a84.1 84.1 0 0 0-78.08-83.77c-9.31 14.09-9.89 27-9.92 27.83a4 4 0 0 0 8-.06a12 12 0 0 1 24 0a28 28 0 0 1-56 0c0-.65.1-11.19 5.52-24.92A84 84 0 1 0 212 128Z"/>`),
		g.Group(children),
	)
}