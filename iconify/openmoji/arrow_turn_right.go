package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="M31.669 28.3c-8.116 0-13.92 5.321-13.92 15.52v18.87h-6.14v-19.2c0-13.5 7.748-20.794 20.248-20.794H51.84l-10.4-9.6l3.8-4.1l17.8 16.5l-17.8 16.5l-3.8-4.1l10.4-9.6Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M31.669 28.3c-8.116 0-13.92 5.321-13.92 15.52v18.87h-6.14v-19.2c0-13.5 7.748-20.794 20.248-20.794H51.84l-10.4-9.6l3.8-4.1l17.8 16.5l-17.8 16.5l-3.8-4.1l10.4-9.6Z"/>`),
		g.Group(children),
	)
}