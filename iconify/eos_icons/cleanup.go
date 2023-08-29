package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cleanup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 15.016h-5v-.013a1.984 1.984 0 0 0-1.978-1.978h-.035l3.857-10.393l-1.732-.643l-4.095 11.036h-.039A1.984 1.984 0 0 0 9 15.003v.013H4v2h16Zm0 3.002H4L2 22h20l-2-3.982z"/>`),
		g.Group(children),
	)
}