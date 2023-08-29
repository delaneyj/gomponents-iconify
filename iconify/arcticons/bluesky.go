package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluesky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.395 28.624a8.482 8.482 0 0 0-8.135-10.915c-.507 0-1.001.052-1.484.138a5.913 5.913 0 0 0-5.747-4.536c-1.983 0-3.724.986-4.795 2.484c-.212-.012-.421-.032-.636-.032c-4.703 0-8.707 2.952-10.288 7.1a5.913 5.913 0 0 0 .105 11.826h29.96a3.125 3.125 0 0 0 3.125-3.125a3.115 3.115 0 0 0-2.106-2.94Z"/>`),
		g.Group(children),
	)
}