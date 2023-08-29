package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFullTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M4.834 7.002A.835.835 0 0 0 4 7.838v4.326c0 .462.374.836.835.836h8.33a.835.835 0 0 0 .836-.836V7.838a.835.835 0 0 0-.835-.836h-8.33zM2 7.554a2.5 2.5 0 0 1 2.5-2.5h9a2.5 2.5 0 0 1 2.5 2.5v.833l1.167.002a.833.833 0 0 1 .833.833v1.667a.833.833 0 0 1-.833.834L16 11.72v.834a2.5 2.5 0 0 1-2.5 2.5h-9a2.5 2.5 0 0 1-2.5-2.5v-5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}