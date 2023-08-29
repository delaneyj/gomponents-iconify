package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Riotgames(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.534 21.77l-1.09-2.81l10.52.54l-.451 4.5zM15.06 0L.307 6.969L2.59 17.471H5.6l-.52-7.512l.461-.144l1.81 7.656h3.126l-.116-9.15l.462-.144l1.582 9.294h3.31l.78-11.053l.462-.144l.82 11.197h4.376l1.54-15.37Z"/>`),
		g.Group(children),
	)
}