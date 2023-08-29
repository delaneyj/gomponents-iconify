package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Syncme(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.247 20.874a21.472 21.472 0 0 0-42.494 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m45.247 20.874l-10.079-4.921L39.1 8.716M2.753 27.126a21.472 21.472 0 0 0 42.494 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m2.753 27.126l10.079 4.921L8.9 39.284"/>`),
		g.Group(children),
	)
}