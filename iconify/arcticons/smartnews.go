package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartnews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm-35 5.89h37M33.25 5.5v5.637M14.75 5.5v5.637M24 5.5v5.637"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.5 16.5h13v10h-13Zm-16 0h13m-13 5h13m0 5h-13m16 5h13m0 5h-13m-16-5h13v5h-13Z"/>`),
		g.Group(children),
	)
}