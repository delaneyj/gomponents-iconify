package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalMilitaryFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M207 32H49a17 17 0 0 0-17 17v49.21a17 17 0 0 0 10 15.47l62.6 28.45a48 48 0 1 0 46.88 0L214 113.68a17 17 0 0 0 10-15.47V49a17 17 0 0 0-17-17ZM96 48h64v72.67l-32 14.54l-32-14.54Zm32 168a32 32 0 1 1 32-32a32 32 0 0 1-32 32Z"/>`),
		g.Group(children),
	)
}