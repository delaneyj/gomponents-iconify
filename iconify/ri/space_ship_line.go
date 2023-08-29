package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceShipLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.88 18.05a35.897 35.897 0 0 1 8.53-16.322a.8.8 0 0 1 1.178 0a35.897 35.897 0 0 1 8.531 16.32a43.99 43.99 0 0 1-6.584.875L12.447 23.1a.5.5 0 0 1-.894 0l-2.089-4.177a44.028 44.028 0 0 1-6.584-.875Zm6.697-1.123l1.158.066L12 19.523l1.265-2.53l1.157-.066a42.139 42.139 0 0 0 4.227-.454a33.914 33.914 0 0 0-6.65-12.387a33.913 33.913 0 0 0-6.648 12.387a42.14 42.14 0 0 0 4.226.454ZM12 14.996a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}