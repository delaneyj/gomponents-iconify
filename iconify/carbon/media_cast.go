package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaCast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 26h3a3 3 0 0 0-3-3zm7 0H7a5.006 5.006 0 0 0-5-5v-2a7.008 7.008 0 0 1 7 7z"/><path fill="currentColor" d="M13 26h-2a9.01 9.01 0 0 0-9-9v-2a11.012 11.012 0 0 1 11 11Z"/><path fill="currentColor" d="M28 26H15v-2h13V8H4v5H2V8a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}