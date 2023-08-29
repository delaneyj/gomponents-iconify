package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stickers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTStickers0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 44c11.046 0 20-8.954 20-20c0 0-10.5 3-17-4S24 4 24 4C12.954 4 4 12.954 4 24s8.954 20 20 20Z"/><path d="M44 24L24 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTStickers0)"/>`),
		g.Group(children),
	)
}