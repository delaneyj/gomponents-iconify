package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Audiofx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m31.304 12.934l7.627-7.617l3.798 3.798l-7.647 7.597Zm5.788 2.158a15.845 15.845 0 0 1-2.087 20.323l3.928 4.069a21.5 21.5 0 0 0 2.224-28.425"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.997 10.974a15.845 15.845 0 0 0-20.002 24.441l-3.928 4.069A21.5 21.5 0 0 1 37.034 6.918"/>`),
		g.Group(children),
	)
}