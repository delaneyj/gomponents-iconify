package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivatTwentyFourBusiness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 5.5v14.8h8.325v-5.55H33.25v19.425h-3.7V42.5H42.5v-37h-37Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.05 29.55H6.425c-.462 0-.925.463-.925.925v11.1c0 .463.463.925.925.925h16.65c.463 0 .925-.463.925-.925v-11.1c0-.463-.463-.925-.925-.925H18.45v-3.7h-7.4v3.7Zm7.4 0h-7.4"/>`),
		g.Group(children),
	)
}