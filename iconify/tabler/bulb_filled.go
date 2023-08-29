package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulbFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M4 11a1 1 0 0 1 .117 1.993L4 13H3a1 1 0 0 1-.117-1.993L3 11h1zm8-9a1 1 0 0 1 .993.883L13 3v1a1 1 0 0 1-1.993.117L11 4V3a1 1 0 0 1 1-1zm9 9a1 1 0 0 1 .117 1.993L21 13h-1a1 1 0 0 1-.117-1.993L20 11h1zM4.893 4.893a1 1 0 0 1 1.32-.083l.094.083l.7.7a1 1 0 0 1-1.32 1.497l-.094-.083l-.7-.7a1 1 0 0 1 0-1.414zm12.8 0a1 1 0 0 1 1.497 1.32l-.083.094l-.7.7a1 1 0 0 1-1.497-1.32l.083-.094l.7-.7zM14 18a1 1 0 0 1 1 1a3 3 0 0 1-6 0a1 1 0 0 1 .883-.993L10 18h4zM12 6a6 6 0 0 1 3.6 10.8a1 1 0 0 1-.471.192L15 17H9a1 1 0 0 1-.6-.2A6 6 0 0 1 12 6z"/></g>`),
		g.Group(children),
	)
}