package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dukto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.614 4.5h8.772c6.199 0 11.19.791 11.19 1.774v14.662c0 .983-4.991 1.774-11.19 1.774h-8.772c-6.199 0-11.19-.791-11.19-1.774V6.274c0-.983 4.991-1.774 11.19-1.774ZM33.99 23.22a.374.374 0 0 1 .003.049v18.427c0 1-4.491 1.804-10.07 1.804s-10.07-.805-10.07-1.804V23.269a.34.34 0 0 1 .004-.048"/>`),
		g.Group(children),
	)
}