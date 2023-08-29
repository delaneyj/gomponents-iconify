package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirusLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12.056 19A6.865 6.865 0 0 0 19 12.056C18.969 8.191 15.81 5.031 11.944 5A6.865 6.865 0 0 0 5 11.944c.031 3.865 3.19 7.025 7.056 7.056ZM18 6l-1.05 1.05M5 5l2 2m10.05 11.05l-.55-.55M6 19.05L7.05 18"/><path stroke="currentColor" stroke-width="1.5" d="M16.5 13a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM11 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><circle cx="9" cy="13" r="1" fill="currentColor"/><circle cx="19.5" cy="4.5" r="1.5" stroke="currentColor" stroke-width="1.5"/><circle cx="1.5" cy="1.5" r="1.5" stroke="currentColor" stroke-width="1.5" transform="matrix(-1 0 0 1 5 2)"/><path stroke="currentColor" stroke-width="1.5" d="M2 12a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0-3 0Z"/><circle cx="1.5" cy="1.5" r="1.5" stroke="currentColor" stroke-width="1.5" transform="matrix(1 0 0 -1 17.05 21.05)"/><circle cx="4.5" cy="20.5" r="1.5" stroke="currentColor" stroke-width="1.5" transform="rotate(180 4.5 20.5)"/><path stroke="currentColor" stroke-width="1.5" d="M13.5 3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6 10a1.5 1.5 0 1 0-.5-2.915v2.83c.156.055.325.085.5.085ZM10.585 19a1.5 1.5 0 0 0 2.83 0h-2.83Z"/></g>`),
		g.Group(children),
	)
}