package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheelchair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1023 1155l102 204q-58 179-210 290t-339 111q-156 0-288.5-77.5t-210-210T0 1184q0-181 104.5-330T379 643l17 131q-122 54-195 165.5T128 1184q0 185 131.5 316.5T576 1632q126 0 232.5-65t165-175.5T1023 1155zm548 100l58 114l-256 128q-13 7-29 7q-40 0-57-35l-239-477H576q-24 0-42.5-16.5T512 935l-96-779q-2-17 6-42q14-51 57-82.5T576 0q66 0 113 47t47 113q0 69-52 117.5T564 319l37 289h423v128H617l16 128h455q40 0 57 35l228 455z"/>`),
		g.Group(children),
	)
}