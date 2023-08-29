package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EcourtsServices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.011 15.194h19.978M15.138 36.152h17.724M24 11.848v24.304m-3.05-7.147c0 1.42-2.47 2.57-5.51 2.57s-5.51-1.15-5.51-2.57c0-.18.04-.36.11-.52l5.4-10.3l5.4 10.3c.07.16.11.34.11.52Zm-10.91-.002h10.8m-5.4-10.818v-2.991m11.61 13.811c0 1.42 2.47 2.57 5.51 2.57s5.51-1.15 5.51-2.57c0-.18-.04-.36-.11-.52l-5.4-10.3l-5.4 10.3c-.07.16-.11.34-.11.52Zm10.91-.002h-10.8m5.4-10.818v-2.991"/>`),
		g.Group(children),
	)
}