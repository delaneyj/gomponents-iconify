package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pexels(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.048 4.5H9.568v39h10.006V33.269h4.474c7.944 0 14.385-6.44 14.385-14.385S31.993 4.5 24.048 4.5Z"/>`),
		g.Group(children),
	)
}