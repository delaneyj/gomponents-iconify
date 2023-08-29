package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareOLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 448v640q0 26-19 45t-45 19q-20 0-37-12L475 820q-27-19-27-52t27-52l448-320q17-12 37-12q26 0 45 19t19 45zm256 800V288q0-13-9.5-22.5T1248 256H288q-13 0-22.5 9.5T256 288v960q0 13 9.5 22.5t22.5 9.5h960q13 0 22.5-9.5t9.5-22.5zm256-960v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}