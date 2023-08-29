package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShovel0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M13 4h22l-1 9l-8.5 6h-3L14 13l-1-9Z"/><path d="M24 19v11"/><rect width="6" height="14" x="21" y="30" fill="#fff" rx="3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShovel0)"/>`),
		g.Group(children),
	)
}