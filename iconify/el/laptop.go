package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M120.264 172.705v628.052h959.473V172.705H120.264zm119.018 114.478h721.436v399.17H239.282v-399.17zM0 857.886v106.201l60.938 63.208h1078.125L1200 964.087V857.886H0zm488.745 34.79h222.51v96.167h-222.51v-96.167z"/>`),
		g.Group(children),
	)
}