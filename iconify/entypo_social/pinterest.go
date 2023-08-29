package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinterest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.617 13.227C8.091 15.981 7.45 18.621 5.549 20c-.586-4.162.861-7.287 1.534-10.605c-1.147-1.93.138-5.812 2.555-4.855c2.975 1.176-2.576 7.172 1.15 7.922c3.891.781 5.479-6.75 3.066-9.199C10.369-.275 3.708 3.18 4.528 8.245c.199 1.238 1.478 1.613.511 3.322c-2.231-.494-2.897-2.254-2.811-4.6c.138-3.84 3.449-6.527 6.771-6.9c4.201-.471 8.144 1.543 8.689 5.494c.613 4.461-1.896 9.293-6.389 8.945c-1.218-.095-1.728-.699-2.682-1.279z"/>`),
		g.Group(children),
	)
}