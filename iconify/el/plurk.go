package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plurk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v1200h230.609V284.605h734.963v652.468H416.6V1200H1200V0z"/><path fill="currentColor" d="M425.599 460.097h353.982v282.206H425.599z"/>`),
		g.Group(children),
	)
}