package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BerberFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#b1cc33" d="M5 30h62v12H5z"/><path fill="#61b2e4" d="M5 17h62v13H5z"/><path fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36 24v24m0-8.437c-6.41 0-12.006 3.393-15 8.437m15-8.437c6.41 0 12.006 3.393 15 8.437M36 32.438c-6.41 0-12.006-3.394-15-8.438m15 8.438c6.41 0 12.006-3.394 15-8.438"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}