package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 3.5l-5.918 10.25L24 24l5.918-10.25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.754 13.75H29.918L24 24h11.836ZM12.164 24H24l-5.918-10.25H6.247ZM24 44.5l5.918-10.25L24 24l-5.918 10.25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.247 34.25h11.835L24 24H12.164ZM35.836 24H24l5.918 10.25h11.835Z"/>`),
		g.Group(children),
	)
}