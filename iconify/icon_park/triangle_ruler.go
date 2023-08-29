package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 44L4 4V44H44Z"/><path fill="#2F88FF" fill-rule="evenodd" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path fill="#2F88FF" fill-rule="evenodd" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path fill="#2F88FF" fill-rule="evenodd" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path fill="#2F88FF" fill-rule="evenodd" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 35H25L13 23V35Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 44V41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 44V41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 44V41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 44V41"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 36H7"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 30H7"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24H7"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 18H7"/></g>`),
		g.Group(children),
	)
}