package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractFifty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 21v135.125h135.125V21H21zm167.438 0v135.125h135.125V21H188.438zm167.437 0v135.125H491V21H355.875zM21 188.438v135.125h135.125V188.438H21zm167.438 0v135.125h135.125V188.438H188.438zm167.437 0v135.125H491V188.438H355.875zM21 355.875V491h135.125V355.875H21zm167.438 0V491h135.125V355.875H188.438zm167.437 0V491H491V355.875H355.875z"/>`),
		g.Group(children),
	)
}