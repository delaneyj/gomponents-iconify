package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prison(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPrison0"><g fill="none" stroke-linecap="round" stroke-width="4"><path stroke="#fff" stroke-linejoin="round" d="M6 4v8a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4M6 7h14m8 15H6v22h22m-12 0V13m-6 9v-9m3-9v3m14 6v3m8-3v3m8-3v3"/><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M43 44V16H27v28h16Z"/><path stroke="#000" d="M35 34v10"/><path stroke="#fff" d="M31 44h8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPrison0)"/>`),
		g.Group(children),
	)
}