package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M17.982 11.969L31.785 4l4.612 7.989l-18.415-.02Z" clip-rule="evenodd"/><path stroke-linejoin="round" d="M4 14a2 2 0 0 1 2-2h36a2 2 0 0 1 2 2v28a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V14Z"/><path stroke-linejoin="round" d="M35.25 33H44V23h-8.75c-2.9 0-5.25 2.239-5.25 5s2.35 5 5.25 5Z"/><path stroke-linecap="round" d="M44 16.5v24"/></g>`),
		g.Group(children),
	)
}