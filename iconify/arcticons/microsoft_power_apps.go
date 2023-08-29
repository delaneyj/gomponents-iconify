package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftPowerApps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.92 8.059L22.433 6.57a2 2 0 0 0-2.93.11L4.852 23.704a2 2 0 0 0 .103 2.719l14.942 14.941a2 2 0 0 0 2.93-.11l1.186-1.379"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.165 15.402l6.175 6.176a2 2 0 0 1 .102 2.718l-7.118 8.273a2 2 0 0 0 .102 2.718l6.142 6.143a2 2 0 0 0 2.93-.11l14.65-17.024a2 2 0 0 0-.103-2.718L28.103 6.636a2 2 0 0 0-2.93.11l-5.11 5.938a2 2 0 0 0 .102 2.718Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.24 28.836l-1.206 1.399a2 2 0 0 1-2.928.108l-6.141-6.141a2 2 0 0 1-.102-2.719l5.11-5.938a2 2 0 0 1 2.93-.11l1.35 1.351"/>`),
		g.Group(children),
	)
}