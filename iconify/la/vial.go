package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.688 3.273L17.27 4.687l1.32 1.32L5.5 19.1c-2 2-2 5.4 0 7.4c1 1 2.4 1.5 3.7 1.5c1.4 0 2.7-.5 3.7-1.5l13.092-13.092l1.25 1.248l1.414-1.414l-9.968-9.969zm1.304 4.135l4.6 4.6L20.6 16h-9.2l8.592-8.592zM9.4 18h9.2l-7.1 7.1c-1.3 1.3-3.3 1.3-4.6 0c-1.3-1.3-1.3-3.3 0-4.6L9.4 18z"/>`),
		g.Group(children),
	)
}