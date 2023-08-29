package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Antarctica(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M130.458 203.435L226.656 58.38L448 107.084v92.378l40.823 49.164l-13.937 107.434l-67.182 85.268L281.328 448L288 352l-64-48l-23.835 80L112 352l-36.342-77.936l13.39-71.775l-54.041-27.343L24.779 112l24.384 49.054L112 176z"/>`),
		g.Group(children),
	)
}