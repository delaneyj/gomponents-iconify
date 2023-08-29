package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 1037.516h1200V1200H0v-162.484zM820.785 0h-441.57v496.632H103.233L600 959.265l496.768-462.633H820.785V0z"/>`),
		g.Group(children),
	)
}