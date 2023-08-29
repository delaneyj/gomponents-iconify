package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualRatio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEqualRatio0"><g fill="none"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="#000" fill-rule="evenodd" d="M24 22.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 17v14m17-14v14"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEqualRatio0)"/>`),
		g.Group(children),
	)
}