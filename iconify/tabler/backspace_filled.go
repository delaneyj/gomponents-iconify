package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M20 5a2 2 0 0 1 1.995 1.85L22 7v10a2 2 0 0 1-1.85 1.995L20 19H9a1 1 0 0 1-.608-.206l-.1-.087l-5.037-5.04c-.809-.904-.847-2.25-.083-3.23l.12-.144l5-5a1 1 0 0 1 .577-.284L9 5h11zm-7.489 4.14a1 1 0 0 0-1.301 1.473l.083.094L12.585 12l-1.292 1.293l-.083.094a1 1 0 0 0 1.403 1.403l.094-.083L14 13.415l1.293 1.292l.094.083a1 1 0 0 0 1.403-1.403l-.083-.094L15.415 12l1.292-1.293l.083-.094a1 1 0 0 0-1.403-1.403l-.094.083L14 10.585l-1.293-1.292l-.094-.083l-.102-.07z"/></g>`),
		g.Group(children),
	)
}