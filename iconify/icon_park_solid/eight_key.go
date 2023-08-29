package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEightKey0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M24 22a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 11a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEightKey0)"/>`),
		g.Group(children),
	)
}