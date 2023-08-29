package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplanemodeInactiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-2.75.775q-.275.075-.513-.1T8.5 21.2v-.4q0-.15.063-.263t.187-.212L10.5 19v-5.5l-7.4 2.175q-.425.125-.763-.138T2 14.85v-.35q0-.225.113-.425t.312-.325L7.8 10.6L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-5.6-5.6V19l1.725 1.275q.125.1.2.25t.075.325v.2q0 .35-.288.562t-.637.113L12 21Zm0-19q.625 0 1.063.438T13.5 3.5V9l8.025 4.725q.225.125.35.338t.125.487v.175q0 .5-.375.775t-.85.15l-3.225-.975L10.5 7.65V3.5q0-.625.438-1.063T12 2Z"/>`),
		g.Group(children),
	)
}