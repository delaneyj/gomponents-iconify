package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartDecreasing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3.087 5.278A3.001 3.001 0 0 1 6 3h20a3 3 0 0 1 3 3v17.247l-6.062-5.985a2.2 2.2 0 0 0-3.06-.03l-2.15 2.041a.5.5 0 0 1-.698-.01L3.087 5.279ZM3 8.114l12.868 12.724a2.2 2.2 0 0 0 3.052.04l2.15-2.017a.5.5 0 0 1 .694.01L29 26.04A3 3 0 0 1 26 29H6a3 3 0 0 1-3-3V8.114Z"/>`),
		g.Group(children),
	)
}