package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BadgeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBadgeTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m16 44l8-4l8 4V24.944A11.955 11.955 0 0 1 24 28a11.955 11.955 0 0 1-8-3.056V44Z"/><path fill="#fff" stroke="#fff" d="M36 16a11.97 11.97 0 0 1-4 8.944A11.955 11.955 0 0 1 24 28a11.955 11.955 0 0 1-8-3.056A11.97 11.97 0 0 1 12 16c0-6.627 5.373-12 12-12s12 5.373 12 12Z"/><path stroke="#000" d="M24 21V11l-2 1m2 9h2m-2 0h-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBadgeTwo0)"/>`),
		g.Group(children),
	)
}