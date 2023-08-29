package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicycleCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><path fill-rule="evenodd" d="M18.5 21.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-7a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm-11 7a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-7a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z" clip-rule="evenodd"/><path d="M8 8a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H8Z"/><path fill-rule="evenodd" d="m13.475 17.343l-.035-.104l3.984-6.355A.5.5 0 0 0 17 10.12h-5.91l-.917-2.776a.5.5 0 1 0-.95.314l1.072 3.245l-3.24 6.371A.5.5 0 0 0 7.5 18H13a.5.5 0 0 0 .475-.657ZM8.315 17h3.993l-1.572-4.76L8.316 17Zm3.104-5.881l1.614 4.887l3.063-4.887H11.42Z" clip-rule="evenodd"/><path d="M16.544 5.999a.5.5 0 1 1-.086-.996C20.113 4.687 22 5.064 22 6.5c0 1.265-.908 1.946-2.483 2a.5.5 0 1 1-.034-1C20.575 7.463 21 7.144 21 6.5c0-.474-1.445-.763-4.457-.502Z"/><path d="m16.993 5.418l2 12c.11.657-.877.822-.986.164l-2-12c-.11-.657.877-.822.986-.164Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}