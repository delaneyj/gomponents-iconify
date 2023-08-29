package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="m70.1 77.57l4.77-8.73l24.52 13.39L84.5 30.06L32.53 45.71l24.55 13.41l-4.8 8.79c-2.56 4.68-7.21 8.37-13.03 8.84c-3.24.27-7.76-.94-10.56-2.48l-.08 21.97c7.62 2.41 28.06 5.93 41.49-18.67z"/>`),
		g.Group(children),
	)
}