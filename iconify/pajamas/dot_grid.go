package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 2.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm11.5 0a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0Zm-1.25 7a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm1.25 4.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM2.25 9.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm7-1.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM8 15a1.25 1.25 0 1 0 0-2.5A1.25 1.25 0 0 0 8 15ZM9.25 2.25a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0ZM2.25 15a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}