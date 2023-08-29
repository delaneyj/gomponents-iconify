package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGiftBag0"><g fill="none" stroke="#fff" stroke-width="4"><rect width="36" height="28" x="6" y="14" stroke-linejoin="round" rx="3"/><path fill="#fff" stroke-linejoin="round" d="M6 32h36v7a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3v-7Z"/><circle cx="19" cy="9" r="5" fill="#fff"/><circle cx="28" cy="10" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="m17 20l7-6l7 6"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGiftBag0)"/>`),
		g.Group(children),
	)
}