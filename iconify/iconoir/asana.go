package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm-5 9a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10 0a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}