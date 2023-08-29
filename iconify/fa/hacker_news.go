package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HackerNews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m809 876l266-499H963L806 689q-24 48-44 92l-42-92l-155-312H445l263 493v324h101V876zm727-588v960q0 119-84.5 203.5T1248 1536H288q-119 0-203.5-84.5T0 1248V288Q0 169 84.5 84.5T288 0h960q119 0 203.5 84.5T1536 288z"/>`),
		g.Group(children),
	)
}