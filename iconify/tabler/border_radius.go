package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRadius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12V8a4 4 0 0 1 4-4h4m4 0v.01M20 4v.01M20 8v.01M20 12v.01M4 16v.01M20 16v.01M4 20v.01M8 20v.01m4-.01v.01m4-.01v.01m4-.01v.01"/>`),
		g.Group(children),
	)
}