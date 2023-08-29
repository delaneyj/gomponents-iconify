package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Small(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M333 83v226c0 10 4 19 11 27c6 7 16 11 27 11h225c32 0 40-17 17-40l-71-73l118-118c3-4 6-9 6-14s-3-9-6-13l-69-69c-4-4-8-6-14-6c-5 0-10 2-14 6L444 139l-71-72c-23-23-40-15-40 16zM5 605l70 69c4 3 8 6 14 6c5 0 9-3 13-6l119-119l72 71c23 23 40 16 40-15V385c0-19-16-38-38-38H70c-32 0-41 17-18 40l72 72L5 577c-3 4-5 10-5 15s2 9 5 13z"/>`),
		g.Group(children),
	)
}