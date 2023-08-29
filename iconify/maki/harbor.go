package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harbor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0C5.5 0 4 1.567 4 3.5a3.492 3.492 0 0 0 2.5 3.338v6.039c-.93-.165-1.875-.55-2.648-1.27c-1.053-.98-1.85-2.54-1.85-5.109a1 1 0 1 0-2.002 0c0 3.003 1.012 5.196 2.49 6.572C3.968 14.447 5.838 15 7.5 15c1.666 0 3.535-.56 5.012-1.94s2.486-3.573 2.486-6.562c.065-1.395-2.063-1.395-1.998 0c0 2.553-.8 4.115-1.853 5.1c-.774.722-1.718 1.11-2.647 1.277V6.842A3.494 3.494 0 0 0 11 3.5C11 1.567 9.5 0 7.5 0zm0 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3z"/>`),
		g.Group(children),
	)
}