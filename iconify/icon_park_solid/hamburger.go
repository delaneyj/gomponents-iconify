package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSHamburger0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" fill-rule="evenodd" d="M44 22c0-9.941-8.954-18-20-18S4 12.059 4 22h40Z" clip-rule="evenodd"/><path fill="#fff" d="M4 38h40v6H4z"/><path d="m4 28l5.455 4l7.272-4L24 32l7.273-4l7.272 4L44 28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSHamburger0)"/>`),
		g.Group(children),
	)
}