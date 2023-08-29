package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardBulletedOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3.35 3.58l17.3 17.29l-1.42 1.42L16.94 20H4a2 2 0 0 1-2-2V6c0-.28.06-.54.16-.78l-.93-.93l1.42-1.42l.7.71M6.6 4H20a2 2 0 0 1 2 2v12c0 .4-.12.77-.32 1.08L17.6 15H20v-2h-4.4l-2-2H20V9h-8.4l-5-5m3.34 9H9v2h2v-.94L9.94 13m-4-4H5v2h2v-.94L5.94 9Z"/>`),
		g.Group(children),
	)
}