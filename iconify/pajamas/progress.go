package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Progress(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 5.5h8a2.5 2.5 0 0 1 0 5H4a2.5 2.5 0 0 1 0-5ZM0 8a4 4 0 0 1 4-4h8a4 4 0 0 1 0 8H4a4 4 0 0 1-4-4Zm4-1a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}