package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 21H6c-1.886 0-2.828 0-3.414-.586C2 19.828 2 18.886 2 17c0-1.886 0-2.828.586-3.414C3.172 13 4.114 13 6 13h12c1.886 0 2.828 0 3.414.586C22 14.172 22 15.114 22 17c0 1.886 0 2.828-.586 3.414C20.828 21 19.886 21 18 21h-1M11 2h7c1.886 0 2.828 0 3.414.586C22 3.172 22 4.114 22 6c0 1.886 0 2.828-.586 3.414C20.828 10 19.886 10 18 10H6c-1.886 0-2.828 0-3.414-.586C2 8.828 2 7.886 2 6c0-1.886 0-2.828.586-3.414C3.172 2 4.114 2 6 2h1m6.5 4H18M6 7V5m3 2V5m4.5 12H18M6 18v-2m3 2v-2"/>`),
		g.Group(children),
	)
}