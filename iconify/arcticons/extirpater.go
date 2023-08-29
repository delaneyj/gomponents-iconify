package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extirpater(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.55 15.694v21.282c0 2.144 2.33 3.87 5.225 3.87h10.449c2.894 0 5.225-1.726 5.225-3.87V15.694Zm-2.075-4.888v1.522a.957.957 0 0 0 .836 1.043h23.378a.957.957 0 0 0 .836-1.043h0v-1.522a.957.957 0 0 0-.836-1.043h-8.348a.957.957 0 0 1-.836-1.043h0v-.522a.957.957 0 0 0-.836-1.043h-3.338a.957.957 0 0 0-.836 1.043h0v.522a.957.957 0 0 1-.836 1.043h-8.348a.957.957 0 0 0-.836 1.043h0Zm9.962 13.596l5.125 7.735m0-7.735l-5.125 7.735"/>`),
		g.Group(children),
	)
}