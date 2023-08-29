package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hsbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.75 13.75L24 24l10.25-10.25h-20.5zm20.5 20.5L24 24L13.75 34.25h20.5zm0 0L44.5 24L34.25 13.75v20.5zm-20.5-20.5L3.5 24l10.25 10.25v-20.5z"/>`),
		g.Group(children),
	)
}