package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 6.5v15m0-15V4m0 2.5H9.5A2.5 2.5 0 1 1 12 4m0 2.5h2.5A2.5 2.5 0 1 0 12 4M3.25 14h17.5m-17.5 0c0-2.328-.23-4.65-.686-6.932L2.5 6.75V6.5h19v.25l-.064.318A35.346 35.346 0 0 0 20.75 14m-17.5 0c0 2.328-.23 4.65-.686 6.932l-.064.318v.25h19v-.25l-.064-.318A35.345 35.345 0 0 1 20.75 14"/>`),
		g.Group(children),
	)
}