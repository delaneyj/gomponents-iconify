package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrightnessUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12a3 3 0 1 0 6 0a3 3 0 1 0-6 0m3-7V3m5 4l1.4-1.4M19 12h2m-4 5l1.4 1.4M12 19v2m-5-4l-1.4 1.4M6 12H4m3-5L5.6 5.6"/>`),
		g.Group(children),
	)
}