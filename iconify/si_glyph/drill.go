package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.999 4a1 1 0 0 0-1-1h-1a1 1 0 0 0-1-1h-1c0-.553-.446-1-.998-1H5.009C1.019 1 .021 3.447.021 4v3C.021 8.983 1.438 9 2 9l-2 5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1s-4 0 0-5h.007C6.559 9 9 9 9 7h1a1 1 0 0 0 1-1h1a1 1 0 0 0 1-1h3V4h-3.001zm-10 0h1v3h-1V4zm-2 0h1v3h-1V4z"/>`),
		g.Group(children),
	)
}