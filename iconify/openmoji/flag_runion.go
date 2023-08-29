package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagRunion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M15 17H5l31 19h0L15 17z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M57 17h10L36 36h0l21-19z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M67 39v-6l-31 3h0l31 3z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M5 39v-6l31 3h0L5 39z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="M33 17h6l-3 19h0l-3-19z"/><path fill="#d22f27" d="M36 36l31 19H5l31-19z"/><g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/></g>`),
		g.Group(children),
	)
}