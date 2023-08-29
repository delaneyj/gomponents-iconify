package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AllFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.59 43.5h-7.051m-4.938 0v-9.804m4.938 5.721v-5.721M17.513 43.5h1.655m4.433-14.026H9.173m29.061 0h-9.695m0-4.469v-12.99m0-4.917V4.5m-4.938 4.918v15.587m-3.685-9.619l-8.958 10.545"/>`),
		g.Group(children),
	)
}