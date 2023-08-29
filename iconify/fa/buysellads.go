package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Buysellads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M915 958H621l147-551zm86 322h311L988 256H548L224 1280h311l383-314zm535-992v960q0 118-85 203t-203 85H288q-118 0-203-85T0 1248V288Q0 170 85 85T288 0h960q118 0 203 85t85 203z"/>`),
		g.Group(children),
	)
}