package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagLesotho(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path fill="#1e50a0" d="M5 17h62v13H5z"/><g stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path d="m36 34.75l-3.464 6h6.928l-3.464-6z"/><ellipse cx="36" cy="34.25" fill="none" rx="2" ry="3"/><path fill="none" d="M36 31.25v6m-1.5-3.5h3"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}