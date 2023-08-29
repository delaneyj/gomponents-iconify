package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicycleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15.5 18.5a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-7a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm-11 7a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-7a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z" clip-rule="evenodd"/><path d="M5 5a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1H5Z"/><path fill-rule="evenodd" d="m10.475 14.343l-.035-.104l3.984-6.355A.5.5 0 0 0 14 7.12H8.09l-.917-2.776a.5.5 0 1 0-.95.314l1.072 3.245l-3.24 6.372A.5.5 0 0 0 4.5 15H10a.5.5 0 0 0 .475-.657ZM5.315 14h3.993L7.736 9.24L5.316 14ZM8.42 8.12l1.614 4.887l3.064-4.887H8.419Z" clip-rule="evenodd"/><path d="M13.544 2.999a.5.5 0 1 1-.086-.996C17.113 1.687 19 2.064 19 3.5c0 1.265-.908 1.946-2.483 2a.5.5 0 1 1-.034-1C17.575 4.463 18 4.144 18 3.5c0-.474-1.445-.763-4.457-.502Z"/><path d="m13.993 2.418l2 12c.11.657-.877.822-.986.164l-2-12c-.11-.657.877-.822.986-.164Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}