package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alpimaps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.419 37.914l10.54-19.222l-4.954-8.606L12.51 37.914h9.908Zm0 0H43.5L32.96 18.692M12.51 37.914H4.5l12.298-22.09l4.039 7.136"/>`),
		g.Group(children),
	)
}