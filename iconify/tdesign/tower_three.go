package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.126V1h2v2.126a4.002 4.002 0 0 1 1.575 6.935L16.127 23H7.873l1.553-12.939A4.002 4.002 0 0 1 11 3.126Zm.334 7.819L10.127 21h3.746l-1.207-10.055a4.023 4.023 0 0 1-1.332 0ZM12 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z"/>`),
		g.Group(children),
	)
}