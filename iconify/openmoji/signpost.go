package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signpost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#6a462f" d="M31.071 10.482h7.947v55.182h-7.947z"/><path fill="#a57939" d="M31.071 18.66v-8.179h7.947v6.077l-7.947 2.102zm11.534-3.668h15.576l4.692 6.18l-4.692 6.249H42.605V14.992z"/><path fill="#6a462f" d="M27.425 27.091H11.848l-4.691 6.18l4.691 6.25h15.577v-12.43zm15.18-3.875l18.37-4.544l1.898 2.5l-4.692 6.249H42.605v-4.205z"/><path fill="#a57939" d="m7.157 33.271l20.268-6.18H11.848l-4.691 6.18z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.23 27.091h-6.382l-4.691 6.18l4.691 6.25h15.577v-5.539m19.286-18.99h11.47l4.692 6.18l-4.692 6.249H42.605v-4.205M31.071 65.663V10.481h7.947v55.182"/>`),
		g.Group(children),
	)
}