package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCardAddThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.5a.5.5 0 0 0 0 1v-1ZM5 3l.499-.038A.5.5 0 0 0 5 2.5V3Zm16 3l.497.055A.5.5 0 0 0 21 5.5V6ZM5.23 6l-.498.038L5.231 6Zm13.109 9.119l.035.498l-.035-.498Zm-10.355.74l-.036-.5l.036.5ZM3 3.5h2v-1H3v1Zm5.02 12.857l10.354-.74l-.071-.997l-10.355.74l.072.997ZM20.68 13.4l.817-7.345l-.994-.11l-.816 7.344l.994.11ZM4.502 3.038l.231 3l.997-.076l-.23-3l-.998.076Zm.231 3l.617 8.017l.997-.077l-.617-8.016l-.997.076ZM21 5.5H5.23v1H21v-1Zm-2.626 10.117a2.5 2.5 0 0 0 2.307-2.217l-.994-.11a1.5 1.5 0 0 1-1.384 1.33l.071.997ZM7.948 15.36a1.5 1.5 0 0 1-1.602-1.382l-.997.077a2.5 2.5 0 0 0 2.67 2.302l-.07-.997Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`),
		g.Group(children),
	)
}