package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNorfolkIsland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="M5 17h62v38H5z"/><path fill="#186648" d="M5 17h21v38H5zm41 0h21v38H46z"/><g stroke="#186648" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="#186648" d="m36 26l-4 4l4-.8l4 .8l-4-4zm0 4l-5 4l5-.8l5 .8l-5-4zm0 4l-6 4l6-.8l6 .8l-6-4zm0 4l-7 4l7-.8l7 .8l-7-4z"/><path fill="none" d="M36 40v6"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}