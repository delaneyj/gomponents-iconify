package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeOneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5 10.494l.257-.429l-.119-.07H3.5v.5Zm0-5.996v.5h.138l.12-.071l-.258-.429Zm5-2.998H9a.5.5 0 0 0-.757-.429L8.5 1.5Zm0 11.992l-.257.429A.5.5 0 0 0 9 13.492h-.5Zm-5-3.498h-2v1h2v-1Zm-2 0a.5.5 0 0 1-.5-.5H0c0 .83.672 1.5 1.5 1.5v-1Zm-.5-.5V5.498H0v3.998h1Zm0-3.997a.5.5 0 0 1 .5-.499v-1a1.5 1.5 0 0 0-1.5 1.5h1Zm.5-.499h2v-1h-2v1Zm2.257-.071l5-2.998l-.514-.858l-5 2.998l.514.858ZM8 1.5v11.992h1V1.5H8Zm.757 11.563l-5-2.998l-.514.858l5 2.998l.514-.858ZM10 6v3h1V6h-1Z"/>`),
		g.Group(children),
	)
}