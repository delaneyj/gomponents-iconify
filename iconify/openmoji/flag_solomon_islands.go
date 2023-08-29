package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSolomonIslands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17v38l62-38H5z"/><path fill="#fcea2b" d="M67 21v-4h-5L5 51v4h5l57-34z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m9.384 25.5l1.66-5l1.431 4.923L8.5 22.457l5-.124L9.384 25.5zm12 0l1.66-5l1.431 4.923l-3.975-2.966l5-.124l-4.116 3.167zm-12 10l1.66-5l1.431 4.923L8.5 32.457l5-.124L9.384 35.5zm6-5l1.66-5l1.431 4.923l-3.975-2.966l5-.124l-4.116 3.167zm6 5l1.66-5l1.431 4.923l-3.975-2.966l5-.124l-4.116 3.167z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}