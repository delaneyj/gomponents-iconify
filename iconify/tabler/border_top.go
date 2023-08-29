package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16M4 8v.01M12 8v.01M20 8v.01M4 12v.01M8 12v.01m4-.01v.01m4-.01v.01m4-.01v.01M4 16v.01m8-.01v.01m8-.01v.01M4 20v.01M8 20v.01m4-.01v.01m4-.01v.01m4-.01v.01"/>`),
		g.Group(children),
	)
}