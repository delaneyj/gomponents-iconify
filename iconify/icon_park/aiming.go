package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aiming(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#2F88FF" fill-rule="evenodd" d="M24 37V44V37Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 37V44"/><path fill="#2F88FF" fill-rule="evenodd" d="M36 24H44H36Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 24H44"/><path fill="#2F88FF" fill-rule="evenodd" d="M4 24H11H4Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 24H11"/><path fill="#2F88FF" fill-rule="evenodd" d="M24 11V4V11Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 11V4"/></g>`),
		g.Group(children),
	)
}