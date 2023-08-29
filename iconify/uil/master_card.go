package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MasterCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.265 5.274a6.681 6.681 0 0 0-3.273.855a6.728 6.728 0 1 0 0 11.745a6.726 6.726 0 1 0 3.273-12.6Zm-5.028 11.183a4.667 4.667 0 0 1-1.518.273a4.728 4.728 0 0 1 0-9.456a4.667 4.667 0 0 1 1.518.273a6.687 6.687 0 0 0 0 8.91Zm1.755-1.057a4.695 4.695 0 0 1 0-6.796a4.695 4.695 0 0 1 0 6.796Zm3.273 1.33a4.667 4.667 0 0 1-1.519-.273a6.687 6.687 0 0 0 0-8.91a4.667 4.667 0 0 1 1.519-.273a4.728 4.728 0 0 1 0 9.456Z"/>`),
		g.Group(children),
	)
}