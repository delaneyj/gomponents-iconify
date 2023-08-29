package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BavariaFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M63.814 18H50.542L66 52.2V54H53.542l-16.27-36H24l16.271 36H27L15 27.45l-9-2.295V36.9L13.729 54H27L6 48.645V36.9l60 15.3V40.455L15 27.45L10.729 18H24l42 10.71v-5.873z"/><path fill="#fff" d="M5 17h62v38H5z"/><path fill="#61B2E4" d="M63.814 18H50.542L66 52.2V54H53.542l-16.27-36H24l16.271 36H27L15 27.45l-9-2.295V36.9L13.729 54H27L6 48.645V36.9l60 15.3V40.455L15 27.45L10.729 18H24l42 10.71v-5.873z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M67 17H5v38h62V17Z"/>`),
		g.Group(children),
	)
}