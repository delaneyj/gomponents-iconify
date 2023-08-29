package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Defibrillator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.55 7.338C-.837 2.742 5.18-.322 7.502 4.274c2.321-4.597 8.339-1.532 5.952 3.064c-.087.167-.203.346-.311.521h-1.808L9.52 5.138a.625.625 0 0 0-1.1.114l-1.648 4.12l-1.33-1.33A.625.625 0 0 0 5 7.86H1.86c-.109-.175-.224-.354-.311-.52zM11 9.11a.626.626 0 0 1-.52-.278L9.139 6.82L7.58 10.717a.625.625 0 0 1-1.022.21L4.741 9.109H2.736a42.67 42.67 0 0 0 4.46 4.674a.464.464 0 0 0 .622 0a43.255 43.255 0 0 0 4.45-4.674z"/>`),
		g.Group(children),
	)
}