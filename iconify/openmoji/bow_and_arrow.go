package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowAndArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#a57939" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.8" d="M51.88 55.586A27.017 27.017 0 0 0 16.384 20.09"/><path fill="#d22f27" d="M31.128 40.842L24.03 47.94v6.852a.602.602 0 0 0 1.029.426l5.893-5.894a.603.603 0 0 0 .176-.425Z"/><path fill="#ea5a47" d="M31.128 40.842L24.03 47.94h-6.852a.602.602 0 0 1-.426-1.028l5.894-5.894a.603.603 0 0 1 .426-.176Z"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m19.228 21.517l31.226 31.225"/><path stroke-miterlimit="10" d="M50.218 57.249a2.352 2.352 0 0 1-.496-2.591a24.668 24.668 0 0 0-32.41-32.41a2.35 2.35 0 0 1-1.857-4.316a29.366 29.366 0 0 1 38.583 38.582a2.353 2.353 0 0 1-3.82.735Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M57.966 14.064L22.158 49.872m30.438-35.808h5.37v5.089"/></g>`),
		g.Group(children),
	)
}