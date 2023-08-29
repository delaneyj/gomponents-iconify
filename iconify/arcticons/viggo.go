package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viggo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.9 38.7c0-16.237-13.163-29.4-29.4-29.4v9.6c10.935 0 19.8 8.865 19.8 19.8h9.6Zm0-29.4h9.6v29.4h-9.6z"/>`),
		g.Group(children),
	)
}