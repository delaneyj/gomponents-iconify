package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForLaos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm0 2c9.61 0 18.104 4.868 23.148 12.266H8.852C13.897 8.868 22.39 4 32 4zm9.834 28c0 5.431-4.402 9.833-9.834 9.833c-5.431 0-9.833-4.402-9.833-9.833S26.57 22.167 32 22.167c5.432-.001 9.834 4.401 9.834 9.833zM32 60c-9.611 0-18.104-4.868-23.148-12.268h46.297C50.104 55.132 41.611 60 32 60z"/>`),
		g.Group(children),
	)
}