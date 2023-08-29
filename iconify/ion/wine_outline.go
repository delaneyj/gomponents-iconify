package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M398.57 80H113.43v16S87.51 272 256 272S398.57 96 398.57 96ZM256 272v160"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M352 432H160"/><path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="32" d="M112 160h288"/>`),
		g.Group(children),
	)
}