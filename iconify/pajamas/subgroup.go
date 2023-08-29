package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Subgroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 4a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm1.5 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm7 4a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-10 5.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}