package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetTemporary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" transform="matrix(0 1 1 0 2.5 2.5)"><path d="M3.987 1.078A8 8 0 1 0 8 0"/><circle cx="8" cy="8" r="2" fill="currentColor"/><path d="M4 5V1H0"/></g>`),
		g.Group(children),
	)
}