package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckField(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M32 96c0-35.3 28.7-64 64-64h224c23.7 0 44.4 12.9 55.4 32h51.8c25.3 0 48.2 14.9 58.5 38l52.8 118.8c.5 1.1.9 2.1 1.3 3.2h4.2c35.3 0 64 28.7 64 64v32c17.7 0 32 14.3 32 32s-14.3 32-32 32h-32c0 53-43 96-96 96s-96-43-96-96H256c0 53-43 96-96 96s-96-43-96-96H32c-17.7 0-32-14.3-32-32s14.3-32 32-32v-32c-17.7 0-32-14.3-32-32v-96c0-17.7 14.3-32 32-32V96zm352 128h85.9l-42.7-96H384v96zM160 432a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm368-48a48 48 0 1 0-96 0a48 48 0 1 0 96 0z"/>`),
		g.Group(children),
	)
}