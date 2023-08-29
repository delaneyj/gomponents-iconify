package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCurvingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#40c0e7" d="M87.63 81.56V66.3l33.03 24.22l-33.03 24.21V99.56h-19c-37.33 0-61.29-10.94-61.29-43.15S31.3 13.27 68.63 13.27h42.75v17.94H68.63c-31.81 0-43.36 6.83-43.36 25.2c0 18.37 11.55 25.21 43.36 25.21h18.99v-.06z"/>`),
		g.Group(children),
	)
}