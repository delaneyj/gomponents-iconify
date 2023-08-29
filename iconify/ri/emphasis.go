package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emphasis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-5.5 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm11 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM18 3v2H8v4h9v2H8v4h10v2H6V3h12Z"/>`),
		g.Group(children),
	)
}