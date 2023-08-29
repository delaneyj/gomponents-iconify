package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><rect width="2" height="5" x="17" y="7.5" rx=".5"/><path d="M4 7.5h3v5H4v-5Zm3.5 0h3v5h-3v-5Zm3.5 0h3v5h-3v-5Z"/><path fill-rule="evenodd" d="M14 4.5H4a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3Zm-11 3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}