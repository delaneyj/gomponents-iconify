package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CactusFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M7 22a1 1 0 0 1-.117-1.993L7 20h2v-6a4 4 0 0 1-3.995-3.8L5 10V9a1 1 0 0 1 1.993-.117L7 9v1a2 2 0 0 0 1.85 1.995L9 12V5a3 3 0 0 1 5.995-.176L15 5v10a2 2 0 0 0 1.995-1.85L17 13V8a1 1 0 0 1 1.993-.117L19 8v5a4 4 0 0 1-3.8 3.995L15 17v3h2a1 1 0 0 1 .117 1.993L17 22H7z"/></g>`),
		g.Group(children),
	)
}