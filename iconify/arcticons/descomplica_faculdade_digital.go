package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DescomplicaFaculdadeDigital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.223 11.212L41.776 5.5v37L6.224 35.83Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.75 21.69a6.069 6.069 0 0 0-6.05-6.05h0a6.069 6.069 0 0 0-6.051 6.05v3.934a6.069 6.069 0 0 0 6.05 6.05h0a6.069 6.069 0 0 0 6.051-6.05m-.002 6.049V7.47"/>`),
		g.Group(children),
	)
}