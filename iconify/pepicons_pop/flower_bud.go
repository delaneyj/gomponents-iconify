package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerBud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16.49 13.713a4.003 4.003 0 0 0 0-7.426A4 4 0 0 0 10 2.877a4 4 0 0 0-6.49 3.41a4.003 4.003 0 0 0 0 7.426a4 4 0 0 0 6.49 3.41a4 4 0 0 0 6.49-3.41Zm-2.108-7.035a1 1 0 0 0 .835 1.334a2 2 0 0 1 0 3.976a1 1 0 0 0-.835 1.334a2 2 0 0 1-3.55 1.783a1 1 0 0 0-1.665 0a2 2 0 0 1-3.55-1.783a1 1 0 0 0-.834-1.333a2 2 0 0 1 0-3.977a1 1 0 0 0 .835-1.334a2 2 0 0 1 3.55-1.783a1 1 0 0 0 1.665 0a2 2 0 0 1 3.55 1.783Z" clip-rule="evenodd"/><path d="M13.5 10a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/></g>`),
		g.Group(children),
	)
}