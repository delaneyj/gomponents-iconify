package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSStretchingOne0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="24" cy="8" r="4" fill="#fff"/><path stroke-linecap="round" stroke-linejoin="round" d="m41 8l-12 9.59V44M10.111 23.25L19 18v10.917L7 41"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSStretchingOne0)"/>`),
		g.Group(children),
	)
}