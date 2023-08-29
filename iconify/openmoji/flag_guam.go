package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagGuam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#92d3f5" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M36.002 46.5A11.736 11.736 0 0 0 42 36a11.736 11.736 0 0 0-6-10.5A11.736 11.736 0 0 0 30 36a11.736 11.736 0 0 0 5.998 10.5Z"/><path fill="#fcea2b" d="M30.64 40a11.242 11.242 0 0 0 5.358 6.5h.004A11.242 11.242 0 0 0 41.36 40Z"/><path fill="none" stroke="#d22f27" stroke-miterlimit="10" stroke-width="2" d="M7 19h58v34H7z"/><path fill="none" stroke="#1e50a0" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M28 39h16"/><path fill="none" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M36.002 46.5A11.736 11.736 0 0 0 42 36a11.736 11.736 0 0 0-6-10.5A11.736 11.736 0 0 0 30 36a11.736 11.736 0 0 0 5.998 10.5Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}