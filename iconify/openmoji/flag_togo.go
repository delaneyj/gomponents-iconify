package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTogo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 32h62v8H5zm0 15h62v8H5zm0-30h62v8H5z"/><path fill="#d22f27" d="M5 17h23v23H5z"/><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m14.415 33l2.655-8l2.289 7.878L13 28.131l8-.198L14.415 33z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}