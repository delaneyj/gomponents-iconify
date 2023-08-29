package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractFortyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M21 21v224.572h179.481v20.856H21V491h111.184V370.27h247.632V491H491V266.43H311.519v-20.857H491V21H379.816v120.731H132.184V21H21zm139.237 0v90.034h191.379V21H160.237zm0 379.966V491h191.379v-90.034H160.237z"/>`),
		g.Group(children),
	)
}