package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Affinitydesigner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m199.842 38.447l-37.922 65.68l15.826 27.404l-78.71 136.328l-47.567 27.465L.033 384.412l.01.002l-.043.023l40.068 69.4h13.914l11.21 19.677h.146l.006-.004l403.22.043l43.424-74.026l.012-.02L310.32 50.192h-23.093l-6.778-11.744h-80.607zm47.465 120.594l123.912 214.09l-145.244-.125l15.418 26.58H108.082L247.307 159.04z"/>`),
		g.Group(children),
	)
}