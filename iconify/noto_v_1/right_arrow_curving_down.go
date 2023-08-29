package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M46.68 21.56L36.22 40.88c3.2 0 7.74 1.1 10.46 2.88c4.89 3.2 7.2 8.67 7.2 14V67.8H25.91l38.11 38.64L102.1 67.8H74.15v-9.96c0-28.02-19.63-34.74-27.47-36.28z"/>`),
		g.Group(children),
	)
}