package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUngroup0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M11.273 4H4v7.273h7.273V4ZM44 36.727h-7.273V44H44v-7.273ZM11.273 24H4v7.273h7.273V24ZM24 36.727h-7.273V44H24v-7.273ZM31.273 4H24v7.273h7.273V4ZM44 16.727h-7.273V24H44v-7.273Z"/><path stroke-linecap="round" d="M11.273 7.636H24m0 32.728h12.727M11.273 27.637h16.364V11.273m1.191 9.09h7.899M20.363 36.727v-9.099v9.099ZM7.637 11.273V24m32.726 0v12.727"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUngroup0)"/>`),
		g.Group(children),
	)
}