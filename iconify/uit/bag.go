package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 6H16V5a2.003 2.003 0 0 0-2-2h-4a2.003 2.003 0 0 0-2 2v1H4.5A2.502 2.502 0 0 0 2 8.5v10A2.502 2.502 0 0 0 4.5 21h15a2.502 2.502 0 0 0 2.5-2.5v-10A2.502 2.502 0 0 0 19.5 6zM9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9V5zm12 13.5a1.5 1.5 0 0 1-1.5 1.5h-15A1.5 1.5 0 0 1 3 18.5v-6.473l5.842 1.948A.51.51 0 0 0 9 14h6a.51.51 0 0 0 .158-.025L21 12.027V18.5zm0-7.494a.49.49 0 0 0-.158.02L14.919 13H9.081l-5.923-1.975a.49.49 0 0 0-.158-.02V8.5A1.5 1.5 0 0 1 4.5 7h15A1.5 1.5 0 0 1 21 8.5v2.506z"/>`),
		g.Group(children),
	)
}