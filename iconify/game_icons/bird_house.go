package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BirdHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 22.93L45.22 203.6l19.31 19.3L256 63.37L447.4 222.9l19.4-19.3zm0 63.86L113 206l54.2 230.1h177.6L399 206zm0 68.31c44.2 0 80 35.8 80 80s-35.8 80-80 80s-80-35.8-80-80s35.8-80 80-80zm0 210c8.8 0 16 7.2 16 16s-7.2 16-16 16s-16-7.2-16-16s7.2-16 16-16zm-9 89v35h18v-35z"/>`),
		g.Group(children),
	)
}