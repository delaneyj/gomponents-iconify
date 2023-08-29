package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowArrowThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m242.83 77.17l-32-32a4 4 0 0 0-5.66 5.66L230.34 76H192a69.84 69.84 0 0 0-26.68 6.37c-12.73 5.88-28.85 18.45-33.27 45c-6.41 38.49-37.53 43.87-48.29 44.57a36 36 0 1 0 0 8a67.53 67.53 0 0 0 22.71-5.54c12.73-5.65 28.86-18.17 33.45-45.71C147.28 84.67 190.18 84 192 84h38.34l-25.17 25.17a4 4 0 0 0 5.66 5.66l32-32a4 4 0 0 0 0-5.66ZM48 204a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/>`),
		g.Group(children),
	)
}