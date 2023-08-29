package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHatMinimalisticBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 10a5 5 0 0 1 5.737-4.946a4.502 4.502 0 0 1 8.526 0A5 5 0 0 1 19 14.584V18c0 1.886 0 2.828-.586 3.414C17.828 22 16.886 22 15 22H9c-1.886 0-2.828 0-3.414-.586C5 20.828 5 19.886 5 18v-3.416A5.001 5.001 0 0 1 2 10Z" opacity=".5"/><path d="M9 17.25a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5H9Z"/></g>`),
		g.Group(children),
	)
}