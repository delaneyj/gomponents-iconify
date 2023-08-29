package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rugby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRugby0"><g fill="none"><g stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" clip-path="url(#ipSRugby1)"><path d="M33.9 33.9c9.372-9.373 12.538-21.403 7.07-26.87c-5.467-5.468-17.497-2.302-26.87 7.07c-9.372 9.373-12.538 21.403-7.07 26.87c5.467 5.468 17.497 2.302 26.87-7.07ZM21.171 21.172l5.657 5.656m-1.414-9.899l5.657 5.657M16.93 25.414l5.656 5.657m-9.894 4.249L35.32 12.692M5.615 28.243l14.142 14.142m8.486-36.77l14.142 14.142"/></g><defs><clipPath id="ipSRugby1"><path fill="#000" d="M0 0h48v48H0z"/></clipPath></defs></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRugby0)"/>`),
		g.Group(children),
	)
}