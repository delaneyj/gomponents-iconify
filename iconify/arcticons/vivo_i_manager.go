package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VivoIManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.693 0 15.274-7.782 15.274-16.966V6.504C35.267 6.503 29.133 4.5 24 4.5S12.723 6.503 8.726 6.503v20.03C8.726 35.719 22.307 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.135 21.997l5.243 5.243l10.487-10.486"/>`),
		g.Group(children),
	)
}