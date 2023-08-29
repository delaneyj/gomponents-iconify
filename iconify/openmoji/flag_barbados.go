package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagBarbados(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f1b31c" d="M5 17h62v38H5z"/><path fill="#1e50a0" d="M5 17h21v38H5zm41 0h21v38H46z"/><g stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path d="M33.415 29.501L36 26.712l2.585 2.789l-.642.595l-1.504-1.629v16.821h-.878V28.467l-1.504 1.629l-.642-.595zm8.085 9.711h-.797V31.9l-1.018 1.103l-.435-.403l2.25-2.388v9zm-11-9l2.25 2.388l-.435.403l-1.018-1.103v7.312H30.5v-9z"/><path d="M41.5 39.297h-11v-.594h11v.594z"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}