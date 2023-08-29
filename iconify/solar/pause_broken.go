package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 18c0 1.886 0 2.828.586 3.414C3.172 22 4.114 22 6 22c1.886 0 2.828 0 3.414-.586C10 20.828 10 19.886 10 18V6c0-1.886 0-2.828-.586-3.414C8.828 2 7.886 2 6 2c-1.886 0-2.828 0-3.414.586C2 3.172 2 4.114 2 6v8m20-8c0-1.886 0-2.828-.586-3.414C20.828 2 19.886 2 18 2c-1.886 0-2.828 0-3.414.586C14 3.172 14 4.114 14 6v12c0 1.886 0 2.828.586 3.414C15.172 22 16.114 22 18 22c1.886 0 2.828 0 3.414-.586C22 20.828 22 19.886 22 18v-8"/>`),
		g.Group(children),
	)
}