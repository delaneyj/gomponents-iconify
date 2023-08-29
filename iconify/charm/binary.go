package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 1.75h3v4.5h-3zm6.5 4.5h3m-3-4.5h1.5v4m-1.5 4h3v4.5h-3zm-6.5 4.5h3m-3-4.5h1.5v4"/>`),
		g.Group(children),
	)
}