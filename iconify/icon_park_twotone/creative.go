package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Creative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCreative0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M42 39V9a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3Z"/><path fill="#555" d="m24 18.316l-2.896 5.91l-6.578.954l4.765 4.658l-1.139 6.478L24 33.199l5.849 3.117l-1.13-6.478l4.756-4.658l-6.541-.954L24 18.316Z"/><path stroke-linecap="round" d="M18.316 12.632h11.368"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCreative0)"/>`),
		g.Group(children),
	)
}