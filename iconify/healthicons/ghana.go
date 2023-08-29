package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghana(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M26 8a2 2 0 1 0-4 0v2.114a13.995 13.995 0 0 0-9.64 6.078A14 14 0 0 0 22 37.827V40a2 2 0 1 0 4 0v-2.173a14 14 0 0 0 7.9-3.957a2 2 0 0 0-2.829-2.829A9.999 9.999 0 0 1 26 33.768V14.172a10 10 0 0 1 5.071 2.727a2 2 0 1 0 2.829-2.828a14 14 0 0 0-7.9-3.957V8Zm-4 6.172a9.996 9.996 0 0 0-6.315 4.243A10 10 0 0 0 22 33.768V14.172Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}