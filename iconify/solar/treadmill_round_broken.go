package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreadmillRoundBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="15" cy="4" r="2" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11 16v-.974c0-.118 0-.177-.002-.234a3 3 0 0 0-.992-2.117a6.004 6.004 0 0 0-.178-.152a5.524 5.524 0 0 1-.252-.216a2 2 0 0 1-.125-2.75c.047-.055.108-.116.231-.239l.33-.33a1.904 1.904 0 0 0-2.356-2.96L4.5 8M3 15.5h.379c1.358 0 2.66-.54 3.621-1.5m5.5-4a4.743 4.743 0 0 0 3 0"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.288 22H3.087a1.087 1.087 0 0 1-.188-2.158l16.157-2.827A2.511 2.511 0 1 1 19.489 22H15"/><path fill="currentColor" d="M18.38 9.91a.75.75 0 1 0 1.483.223L18.38 9.91Zm3.767-2.175a.75.75 0 1 0-.294-1.47l.294 1.47Zm-2.844 6.137a.75.75 0 0 0-1.484-.223l1.484.223Zm.56-3.74a2.88 2.88 0 0 1 2.284-2.397l-.294-1.47A4.38 4.38 0 0 0 18.38 9.91l1.483.223Zm-1.121 7.48l.56-3.74l-1.483-.223l-.56 3.74l1.483.222Z"/></g>`),
		g.Group(children),
	)
}