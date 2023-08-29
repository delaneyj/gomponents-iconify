package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeSlashes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSThreeSlashes0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m6 10l36-6v7L6 17v-7Zm0 14l36-6v7L6 31v-7Zm0 14l36-6v7L6 45v-7Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSThreeSlashes0)"/>`),
		g.Group(children),
	)
}