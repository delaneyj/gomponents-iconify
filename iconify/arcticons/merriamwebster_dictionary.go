package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MerriamwebsterDictionary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.73 4.5H10.4a2 2 0 0 0-2 2v35a2 2 0 0 0 2 2h2.33m0-39v39H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2H12.73Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.965 25.509l-2.4 8l-2.4-8l-2.4 8l-2.4-8m-1 .012h2m2.8 0h2m2.8 0h2m-10.8-3.039l1-7.991l4 8l4-7.988l1 7.988m-11-.009h2m8 0h2m-2-7.979h1m-10 0h1"/>`),
		g.Group(children),
	)
}