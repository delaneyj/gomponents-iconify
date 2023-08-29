package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeSquaresBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><rect width="10" height="10" x="12" y="2" rx="2"/><path d="M12 7h-1c-1.886 0-2.828 0-3.414.586C7 8.172 7 9.114 7 11v2c0 1.886 0 2.828.586 3.414C8.172 17 9.114 17 11 17h2c1.886 0 2.828 0 3.414-.586C17 15.828 17 14.886 17 13v-1h-1c-1.886 0-2.828 0-3.414-.586C12 10.828 12 9.886 12 8V7Z" opacity=".7"/><path d="M7 12v1c0 1.886 0 2.828.586 3.414C8.172 17 9.114 17 11 17h1v1c0 1.886 0 2.828-.586 3.414C10.828 22 9.886 22 8 22H6c-1.886 0-2.828 0-3.414-.586C2 20.828 2 19.886 2 18v-2c0-1.886 0-2.828.586-3.414C3.172 12 4.114 12 6 12h1Z" opacity=".4"/></g>`),
		g.Group(children),
	)
}