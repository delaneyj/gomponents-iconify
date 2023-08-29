package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RefreshCircular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.583 9.667C15.81 8.097 14.043 7 11.988 7C9.388 7 7.25 8.754 7 11"/><path stroke-linecap="round" stroke-linejoin="round" d="M14.494 9.722H16.4a.6.6 0 0 0 .6-.6V7.5m-9.583 6.167C8.191 15.629 9.957 17 12.012 17c2.6 0 4.736-2.193 4.988-5"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.506 13.622H7.6a.6.6 0 0 0-.6.6V16.4"/></g>`),
		g.Group(children),
	)
}