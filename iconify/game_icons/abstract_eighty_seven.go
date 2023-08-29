package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractEightySeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 21v71.675h71.675V21H21zm104.575 0v71.675h260.85V21h-260.85zm293.75 0v71.675H491V21h-71.675zM21 125.575v260.85h71.675v-260.85H21zm104.575 0v260.85h31.872c2.667-45.877 36.674-83.249 80.928-91.21v-78.43c-44.254-7.961-78.261-45.333-80.928-91.21h-31.872zm228.978 0c-2.667 45.877-36.674 83.249-80.928 91.21v78.43c44.254 7.961 78.261 45.333 80.928 91.21h31.872v-260.85h-31.872zm64.772 0v260.85H491v-260.85h-71.675zM21 419.325V491h71.675v-71.675H21zm104.575 0V491h260.85v-71.675h-260.85zm293.75 0V491H491v-71.675h-71.675z"/>`),
		g.Group(children),
	)
}