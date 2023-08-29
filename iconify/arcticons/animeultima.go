package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animeultima(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.375 18.167v10.04l7.25-5.02Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.938 43.5l14.343-39h7.438l14.344 39h-8.5l-2.657-7.905H17.094L14.438 43.5Z"/>`),
		g.Group(children),
	)
}