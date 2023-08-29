package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightingCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18Zm11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-10.915 7.5H10v-5H6.29l5.625-10H14v5h3.71l-5.625 10ZM12 15.571l2.29-4.071H12V8.429L9.71 12.5H12v3.071Z"/>`),
		g.Group(children),
	)
}