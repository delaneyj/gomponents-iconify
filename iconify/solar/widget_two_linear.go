package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetTwoLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M2.5 6.5a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm11 11a4 4 0 1 1 8 0a4 4 0 0 1-8 0Zm8-11c0-1.886 0-2.828-.586-3.414C20.328 2.5 19.386 2.5 17.5 2.5c-1.886 0-2.828 0-3.414.586c-.586.586-.586 1.528-.586 3.414c0 1.886 0 2.828.586 3.414c.586.586 1.528.586 3.414.586c1.886 0 2.828 0 3.414-.586c.586-.586.586-1.528.586-3.414Zm-11 11c0-1.886 0-2.828-.586-3.414C9.328 13.5 8.386 13.5 6.5 13.5c-1.886 0-2.828 0-3.414.586c-.586.586-.586 1.528-.586 3.414c0 1.886 0 2.828.586 3.414c.586.586 1.528.586 3.414.586c1.886 0 2.828 0 3.414-.586c.586-.586.586-1.528.586-3.414Z"/>`),
		g.Group(children),
	)
}