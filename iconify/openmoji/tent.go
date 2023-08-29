package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path fill="#FCEA2B" d="M36 18L14.54 55h42.92z"/><path fill="#F1B31C" d="M36 18L14.54 55H36z"/><path fill="#8967AA" d="M36 46.12c-1.48 0-4.44 1.48-7.4 1.48a7.4 7.4 0 0 0-7.4 7.4H36v-8.88z"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M36 18L14.54 55h42.92z"/><path d="M36 18L14.54 55H36zm0 37V18"/><path d="M36 46.12c-1.48 0-4.44 1.48-7.4 1.48a7.4 7.4 0 0 0-7.4 7.4H36v-8.88zM36 55V18"/></g>`),
		g.Group(children),
	)
}