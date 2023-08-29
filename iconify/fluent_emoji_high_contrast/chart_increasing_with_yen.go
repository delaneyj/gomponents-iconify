package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartIncreasingWithYen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 2a4 4 0 0 0-4 4v20a4 4 0 0 0 4 4h20a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4H6Zm.475 3.43l2.068 3.898l2.206-3.915a.75.75 0 0 1 1.307.737l-2.402 4.26h1.748a.75.75 0 1 1 0 1.5h-2.13v.95h2.13a.75.75 0 0 1 0 1.5h-2.13v1.4a.75.75 0 0 1-1.5 0v-1.4h-1.82a.75.75 0 0 1 0-1.5h1.82v-.95h-1.82a.75.75 0 0 1 0-1.5H7.42L5.15 6.132a.75.75 0 1 1 1.326-.703Zm20.314 7.955a.75.75 0 0 1 0 1.06L16.365 24.87a1.75 1.75 0 0 1-2.48-.005l-2.617-2.637a.25.25 0 0 0-.354 0l-4.422 4.421a.75.75 0 0 1-1.06-1.06l4.422-4.422a1.75 1.75 0 0 1 2.48.005l2.616 2.637a.25.25 0 0 0 .355 0l10.423-10.423a.75.75 0 0 1 1.061 0Z"/>`),
		g.Group(children),
	)
}