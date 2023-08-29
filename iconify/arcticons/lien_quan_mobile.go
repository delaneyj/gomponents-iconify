package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LienQuanMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.676 24c0-.277 3.047-7.203 3.324-7.203s3.324 6.926 3.324 7.203s-3.047 6.649-3.324 6.649s-3.324-6.372-3.324-6.649Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.338 7.932l-6.095 14.683L9.595 24l6.648 1.385L24.554 44.5M24 3.5l7.757 19.115L38.405 24l-6.648 1.385l-6.095 14.682"/>`),
		g.Group(children),
	)
}