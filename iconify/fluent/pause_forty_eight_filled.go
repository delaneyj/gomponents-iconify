package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.75 6A3.75 3.75 0 0 0 8 9.75v28.5A3.75 3.75 0 0 0 11.75 42h6.5A3.75 3.75 0 0 0 22 38.25V9.75A3.75 3.75 0 0 0 18.25 6h-6.5Zm18 0A3.75 3.75 0 0 0 26 9.75v28.5A3.75 3.75 0 0 0 29.75 42h6.5A3.75 3.75 0 0 0 40 38.25V9.75A3.75 3.75 0 0 0 36.25 6h-6.5Z"/>`),
		g.Group(children),
	)
}