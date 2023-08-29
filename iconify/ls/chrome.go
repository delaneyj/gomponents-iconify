package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chrome(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M680 200H358c-31 0-59 10-84 25c-43 27-72 75-74 129L76 139C142 55 243 0 358 0c141 0 264 82 322 200zM60 159l160 278c1 1 1 2 2 3c14 25 36 45 62 59c23 11 47 19 74 19s52-8 75-19L308 714C134 689 0 540 0 359c0-74 22-143 60-200zm382 66h249c16 41 26 87 26 134c0 198-161 358-359 358c-8 0-15 0-23-1l160-276c1-1 1-2 2-3c13-24 20-50 20-78v-5c-2-54-32-102-75-129zM239 359c0-67 52-120 119-120s120 53 120 120s-53 120-120 120s-119-53-119-120z"/>`),
		g.Group(children),
	)
}