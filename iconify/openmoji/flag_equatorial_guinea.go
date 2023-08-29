package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagEquatorialGuinea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path fill="#5c9e31" d="M5 17h62v13H5z"/><path fill="#1e50a0" d="M26 36L5 55V17l21 19z"/><path fill="none" stroke="#6a462f" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M40 34.5v4"/><path fill="#5c9e31" stroke="#5c9e31" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M38 33.5h4"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}