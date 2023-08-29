package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotThirtySix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 37.5v-27a2 2 0 0 0-2-2h-35a2 2 0 0 0-2 2v27a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 34.9V13c0-.884-.761-1.6-1.7-1.6H9.2c-.939 0-1.7.716-1.7 1.6h0v22c0 .884.761 1.6 1.7 1.6h29.6c.939 0 1.7-.716 1.7-1.6v-.1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 26l2-10l.5 14l2-8l3 3l2-10l.5 14.5l2-6.5l3 2l1-3l.8 9l1.2-2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.5 15.9h14V29h-12v-3H31v-4h-7.5v-6.1ZM18 32v1.5m7.5-1.5v1.5h-15V32M31 26h6.5M31 22v-6.1"/>`),
		g.Group(children),
	)
}