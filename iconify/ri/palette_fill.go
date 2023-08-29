package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.522 0 10 3.978 10 8.889a5.558 5.558 0 0 1-5.556 5.555h-1.966c-.922 0-1.667.745-1.667 1.667c0 .422.167.811.422 1.1c.267.3.434.689.434 1.122C13.667 21.256 12.9 22 12 22C6.478 22 2 17.522 2 12S6.478 2 12 2ZM7.5 12a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm9 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM12 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}