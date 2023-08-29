package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m18 18l.01-12.28c.59-.35.99-.99.99-1.72c0-1.1-.9-2-2-2s-2 .9-2 2c0 .8.47 1.48 1.14 1.8l-4.13 6.58c-.33-.24-.73-.38-1.16-.38c-.84 0-1.55.51-1.85 1.24l-2.14-1.53c.09-.22.14-.46.14-.71c0-1.11-.89-2-2-2a2 2 0 0 0-2 2a2 2 0 0 0 .98 1.71L1 18h17zM17 3c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1zM5 10c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1zm5.85 3c.55 0 1 .45 1 1s-.45 1-1 1s-1-.45-1-1s.45-1 1-1z"/>`),
		g.Group(children),
	)
}