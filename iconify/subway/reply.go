package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M185.2 128.6V19.7L0 204.9l185.2 185.2v-109c152.5 0 250.5 0 326.8 217.9c0-108.9 10.9-370.4-326.8-370.4z"/>`),
		g.Group(children),
	)
}