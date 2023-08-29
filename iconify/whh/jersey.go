package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jersey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M769 889q0 23-18 46.5T707 969q-155 55-323 55T61 969q-25-10-43-33.5T0 889V519q0-45 10-88.5t22-70t22-55T64 256V0h128v64q0 80 56.5 136t136 56T520 200t56-136V0h129q-1 169 0 256q0 21 10 49.5t22 55t22 70t10 88.5v370zM320 480q0-13-9-22.5t-23-9.5h-64q-13 0-22.5 9.5t-9.5 23t9.5 22.5t22.5 9h32v225q0 13 9.5 22.5t23 9.5t22.5-9.5t9-22.5V480zm192 0q0-13-9-22.5t-23-9.5h-64q-13 0-22.5 9.5t-9.5 23t9.5 22.5t22.5 9h32v225q0 13 9.5 22.5t23 9.5t22.5-9.5t9-22.5V480z"/>`),
		g.Group(children),
	)
}