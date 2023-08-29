package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemixAddThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 3a1 1 0 0 1 1-1h13c.808 0 1.6.068 2.37.2C24.972 3.325 30 9.075 30 16c0 4.181-1.833 7.934-4.74 10.5h-3.446A11.997 11.997 0 0 0 28 16a11.962 11.962 0 0 0-1.096-5.018C25.004 6.861 20.836 4 16 4H3a1 1 0 0 1-1-1Zm0 13a13.954 13.954 0 0 0 1.279 5.853C5.495 26.663 10.357 30 16 30h13a1 1 0 1 0 0-2H16c-.75 0-1.483-.069-2.194-.2C8.226 26.77 4 21.878 4 16c0-4.518 2.497-8.453 6.186-10.5H6.74A13.966 13.966 0 0 0 2 16Zm14-6a1 1 0 0 1 1 1v4h4a1 1 0 1 1 0 2h-4v4a1 1 0 1 1-2 0v-4h-4a1 1 0 1 1 0-2h4v-4a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}