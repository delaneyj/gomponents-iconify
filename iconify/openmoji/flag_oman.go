package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 42h62v13H5z"/><path fill="#fff" d="M5 17h62v13H5z"/><path fill="#d22f27" d="M5 17h16v38H5z"/><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round"><path d="M10.5 23.917h5m-2.5 0v-4m4 .166l-4 4s-2 4-5 4"/><path d="m9 20.083l4 4s2 4 5 4"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}