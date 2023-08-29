package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyVodafone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 1 2.5 24A21.51 21.51 0 0 1 24 2.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.79 13.41a3.27 3.27 0 0 1 .31-1.21c1.09-2.54 4.9-6.29 7.77-7.29v-.08c-13.47 0-20.6 8.35-20.6 19.17a10.75 10.75 0 1 0 12.52-10.59Z"/>`),
		g.Group(children),
	)
}