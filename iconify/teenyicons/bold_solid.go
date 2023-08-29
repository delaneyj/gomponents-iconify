package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoldSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 1.58c0-.32.26-.58.58-.58H8.5a3.5 3.5 0 0 1 2.21 6.215A3.501 3.501 0 0 1 9.5 14H3.59a.59.59 0 0 1-.59-.59V1.58ZM4 8v5h5.5a2.5 2.5 0 0 0 0-5H4Zm0-1h4.5a2.5 2.5 0 0 0 0-5H4v5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}