package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Form(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="32" x="4" y="8" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4" rx="2"/><path fill="#2F88FF" fill-rule="evenodd" d="M4 29H44H4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 29H44"/><path fill="#2F88FF" fill-rule="evenodd" d="M4 19H44H4Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 19H44"/><path fill="#2F88FF" fill-rule="evenodd" d="M17 40V19V40Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M17 40V19"/><path fill="#2F88FF" fill-rule="evenodd" d="M4 38V17V38Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 38V17"/><path fill="#2F88FF" fill-rule="evenodd" d="M44 38V17V38Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 38V17"/><path fill="#2F88FF" fill-rule="evenodd" d="M31 40V19V40Z" clip-rule="evenodd"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 40V19"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M9 40H39"/></g>`),
		g.Group(children),
	)
}