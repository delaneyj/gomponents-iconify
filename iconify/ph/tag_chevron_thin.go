package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagChevronThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m234 121.34l-42.67-64a12 12 0 0 0-10-5.34H24a4 4 0 0 0-3.29 6.27l46.61 67.51a4 4 0 0 1 0 4.39l-46.61 67.56A4 4 0 0 0 24 204h157.33a12 12 0 0 0 10-5.34l42.67-64a12 12 0 0 0 0-13.32Zm-6.66 8.88l-42.66 64a4 4 0 0 1-3.33 1.78H31.62L74 134.66a12 12 0 0 0 0-13.37L31.62 60h149.71a4 4 0 0 1 3.33 1.78l42.66 64a4 4 0 0 1 0 4.44Z"/>`),
		g.Group(children),
	)
}