package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBosniaAndHerzegovina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M50 55L15 17h35v38z"/><circle cx="11.1" cy="18" r="1.75" fill="#fff"/><circle cx="44.1" cy="54" r="1.75" fill="#fff"/><circle cx="40.1" cy="49.5" r="1.75" fill="#fff"/><circle cx="35.1" cy="45" r="1.75" fill="#fff"/><circle cx="31.1" cy="40.5" r="1.75" fill="#fff"/><circle cx="23.1" cy="31.5" r="1.75" fill="#fff"/><circle cx="27.1" cy="36" r="1.75" fill="#fff"/><circle cx="15.1" cy="22.5" r="1.75" fill="#fff"/><circle cx="19.1" cy="27" r="1.75" fill="#fff"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}