package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm10.392-10a4 4 0 1 0 6.929-4a4 4 0 0 0-6.928 4Zm0 12a4 4 0 1 0 6.929 4a4 4 0 0 0-6.928-4ZM24 36a4 4 0 1 0 0 8a4 4 0 0 0 0-8Zm-10.392-6a4 4 0 1 0-6.928 4a4 4 0 0 0 6.928-4Zm0-12a4 4 0 1 0-6.928-4a4 4 0 0 0 6.928 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}