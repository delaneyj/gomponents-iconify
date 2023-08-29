package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFontSearch0"><g fill="none" stroke-linecap="round" stroke-width="4"><circle cx="22.834" cy="22.834" r="17" fill="#fff" stroke="#fff" stroke-linejoin="round"/><path stroke="#fff" d="m35 35l6 6"/><path stroke="#000" stroke-linejoin="round" d="M23 17v14m-5-14h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFontSearch0)"/>`),
		g.Group(children),
	)
}