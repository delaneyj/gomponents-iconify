package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseBargainButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="48.984" height="48.984" x="11.792" y="11.974" fill="#d0cfce" rx="1.699"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M34.293 24.315h14.135v10.826H34.047V24.315h.246z"/><path stroke-miterlimit="10" d="M34.047 29.728h14.381"/><path stroke-linecap="round" stroke-miterlimit="10" d="M31.693 40.464h19.091m-19.091 5.323h19.091M34.715 51.11h3.533M21.245 36.317l7.476-3.176m-7.795-4.495l7.476-3.176"/><path stroke-linecap="round" stroke-linejoin="round" d="M42.521 52.493h3.533V40.696"/><path stroke-linecap="round" stroke-miterlimit="10" d="M25.465 34.525v17.968"/><rect width="48.984" height="48.984" x="11.792" y="11.974" stroke-miterlimit="10" rx="1.699"/></g>`),
		g.Group(children),
	)
}