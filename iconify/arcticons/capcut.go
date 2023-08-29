package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capcut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 35.9L4.505 16.241V14.86a3.717 3.717 0 0 1 3.722-3.728h21.997a3.717 3.717 0 0 1 3.723 3.728v1.695"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 11.835l-39 19.66l.005 1.646a3.717 3.717 0 0 0 3.723 3.728h21.996a3.717 3.717 0 0 0 3.723-3.728v-1.948"/>`),
		g.Group(children),
	)
}