package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17c-2.828 0-4.243 0-5.121-.879c-.57-.569-.77-1.363-.84-2.621h19.923c-.07 1.258-.271 2.052-.84 2.621C20.241 17 18.827 17 16 17h-3.25v4H16a.75.75 0 0 1 0 1.5H8A.75.75 0 0 1 8 21h3.25v-4H8Zm2-15h4c3.771 0 5.657 0 6.828 1.172C22 4.343 22 6.229 22 10v1c0 .552 0 1.05-.006 1.5H2.007C2 12.05 2 11.552 2 11v-1c0-3.771 0-5.657 1.172-6.828C4.343 2 6.229 2 10 2Z"/>`),
		g.Group(children),
	)
}