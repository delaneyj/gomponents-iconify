package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.463 2.057a.75.75 0 0 1 .817.163l6 6a.75.75 0 0 1-.036 1.094L14.889 14l5.355 4.686a.75.75 0 0 1 .036 1.094l-6 6a.75.75 0 0 1-1.28-.53v-9.597l-4.756 4.161a.75.75 0 1 1-.988-1.128L12.611 14L7.256 9.314a.75.75 0 1 1 .988-1.128L13 12.347V2.75a.75.75 0 0 1 .463-.693ZM14.5 15.653v7.786l4.153-4.152l-4.153-3.634Zm0-3.306l4.153-3.634L14.5 4.561v7.786Z"/>`),
		g.Group(children),
	)
}