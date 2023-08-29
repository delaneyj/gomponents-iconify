package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampfireFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M19.757 16.03a1 1 0 0 1 .597 1.905l-.111.035l-16 4a1 1 0 0 1-.597-1.905l.111-.035l16-4z"/><path fill="currentColor" d="M3.03 16.757a1 1 0 0 1 1.098-.749l.115.022l16 4a1 1 0 0 1-.37 1.962l-.116-.022l-16-4a1 1 0 0 1-.727-1.213zM13.553 2.106C9.379 4.192 7 7.464 7 11a5 5 0 0 0 10 0c0-1.047-.188-1.808-.606-2.705l-.169-.345l-.33-.647C15.274 6.063 15 4.965 15 3a1 1 0 0 0-1.447-.894z"/></g>`),
		g.Group(children),
	)
}