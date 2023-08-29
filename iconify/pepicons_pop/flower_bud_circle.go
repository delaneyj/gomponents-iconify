package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerBudCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M19.49 16.713a4.003 4.003 0 0 0 0-7.426A4 4 0 0 0 13 5.877a4 4 0 0 0-6.49 3.41a4.003 4.003 0 0 0 0 7.427a4 4 0 0 0 6.49 3.41a4 4 0 0 0 6.49-3.41Zm-2.108-7.035a1 1 0 0 0 .835 1.334a2 2 0 0 1 0 3.976a1 1 0 0 0-.835 1.334a2 2 0 0 1-3.55 1.783a1 1 0 0 0-1.665 0a2 2 0 0 1-3.55-1.783a1 1 0 0 0-.834-1.333a2 2 0 0 1 0-3.977a1 1 0 0 0 .835-1.334a2 2 0 0 1 3.55-1.783a1 1 0 0 0 1.665 0a2 2 0 0 1 3.55 1.783Z" clip-rule="evenodd"/><path d="M16.5 13a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}