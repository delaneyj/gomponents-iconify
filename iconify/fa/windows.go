package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M682 878v651L0 1435V878h682zm0-743v659H0V229zm982 743v786l-907-125V878h907zm0-878v794H757V125z"/>`),
		g.Group(children),
	)
}