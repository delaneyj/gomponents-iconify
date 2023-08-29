package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEightTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.834 6.949A.835.835 0 0 0 4 7.784v4.327c0 .461.374.835.835.835h6.33a.835.835 0 0 0 .836-.835V7.784a.835.835 0 0 0-.835-.835h-6.33ZM2 7.5A2.5 2.5 0 0 1 4.5 5h9A2.5 2.5 0 0 1 16 7.5v.833l1.167.003a.833.833 0 0 1 .833.833v1.667a.833.833 0 0 1-.833.833L16 11.667v.833a2.5 2.5 0 0 1-2.5 2.5h-9A2.5 2.5 0 0 1 2 12.5v-5Z"/>`),
		g.Group(children),
	)
}