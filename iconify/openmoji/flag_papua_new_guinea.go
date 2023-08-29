package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPapuaNewGuinea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path d="M5 55V17l62 38H5z"/><circle cx="19" cy="50" r="1" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="19" cy="32" r="1" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="40" r="1" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="26" cy="40" r="1" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="21.5" cy="44.5" r=".5" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M42 25c9 1 15 6.087 15 11a5.166 5.166 0 0 1-3 5"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M50 21c-1.75.658-3.154 3.571-1.9 5l1.9 1Zm0 6c5-3 9 0 9 3h-5l-4-3m-5.19 4.19C44 28 46 26 47.81 26.29l2 .9Zm5-4c-.81 3.81-1 8 2 8l2-5l-4-3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}