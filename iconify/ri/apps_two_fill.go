package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 11.5a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0 10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm10-10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Zm0 10a4.5 4.5 0 1 1 0-9a4.5 4.5 0 0 1 0 9Z"/>`),
		g.Group(children),
	)
}