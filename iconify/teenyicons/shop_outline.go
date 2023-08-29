package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 2.5V2a.5.5 0 0 0-.49.402l.49.098Zm12 0l.49-.098A.5.5 0 0 0 13.5 2v.5Zm1 5V8a.5.5 0 0 0 .49-.598l-.49.098Zm-14 0l-.49-.098A.5.5 0 0 0 .5 8v-.5Zm3 3H3v.5h.5v-.5Zm8 0v.5h.5v-.5h-.5ZM0 15h15v-1H0v1Zm1-7.5v7h1v-7H1Zm12 0v7h1v-7h-1ZM1.5 3h12V2h-12v1Zm11.51-.402l1 5l.98-.196l-1-5l-.98.196ZM14.5 7H.5v1h14V7ZM.99 7.598l1-5l-.98-.196l-1 5l.98.196ZM1 1h13V0H1v1Zm2 6.5v3h1v-3H3Zm.5 3.5h8v-1h-8v1Zm8.5-.5v-3h-1v3h1Z"/>`),
		g.Group(children),
	)
}