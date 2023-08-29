package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSeychelles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#d22f27" d="M67.253 29.579L67 17l-19.839-.055L6 54Z"/><path fill="#f1b31c" d="M47.13 16.98H26.56L5.99 54z"/><path fill="#fff" d="m67.2 41.76l-.04-12.23L6 53.99z"/><path fill="#5c9e31" d="M66.998 54.985V41.78l-61 12.2z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}