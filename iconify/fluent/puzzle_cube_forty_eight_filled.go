package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzleCubeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 6v10.5h10V6H19Zm-2.5 13H6v10h10.5V19ZM19 29V19h10v10H19Zm-2.5 2.5H6v4.25A6.25 6.25 0 0 0 12.25 42h4.25V31.5ZM19 42h10V31.5H19V42Zm12.5 0V31.5H42v4.25A6.25 6.25 0 0 1 35.75 42H31.5ZM42 19v10H31.5V19H42Zm0-2.5v-4.25A6.25 6.25 0 0 0 35.75 6H31.5v10.5H42Z"/>`),
		g.Group(children),
	)
}