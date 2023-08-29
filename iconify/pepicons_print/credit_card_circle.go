package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M8.5 9h12a2.5 2.5 0 0 1 2.5 2.5v7a2.5 2.5 0 0 1-2.5 2.5h-12A2.5 2.5 0 0 1 6 18.5v-7A2.5 2.5 0 0 1 8.5 9Z" opacity=".2"/><path fill-rule="evenodd" d="M19 7H7a2.5 2.5 0 0 0-2.5 2.5v7A2.5 2.5 0 0 0 7 19h12a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 19 7ZM5.5 9.5A1.5 1.5 0 0 1 7 8h12a1.5 1.5 0 0 1 1.5 1.5v7A1.5 1.5 0 0 1 19 18H7a1.5 1.5 0 0 1-1.5-1.5v-7Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 9.5h15a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-15A.5.5 0 0 1 5 11v-1a.5.5 0 0 1 .5-.5Zm3.5 3H8a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1Zm-1 2v-1h1v1H8Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}