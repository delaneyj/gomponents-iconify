package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WeixinCardsOffers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSWeixinCardsOffers0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="m6 12l36 4v24L6 36V12Z"/><path d="M38 15.555V8L6 12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSWeixinCardsOffers0)"/>`),
		g.Group(children),
	)
}