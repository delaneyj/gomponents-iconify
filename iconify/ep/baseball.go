package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Baseball(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M195.2 828.8a448 448 0 1 1 633.6-633.6a448 448 0 0 1-633.6 633.6zm45.248-45.248a384 384 0 1 0 543.104-543.104a384 384 0 0 0-543.104 543.104z"/><path fill="currentColor" d="M497.472 96.896c22.784 4.672 44.416 9.472 64.896 14.528a256.128 256.128 0 0 0 350.208 350.208c5.056 20.48 9.856 42.112 14.528 64.896A320.128 320.128 0 0 1 497.472 96.896zM108.48 491.904a320.128 320.128 0 0 1 423.616 423.68c-23.04-3.648-44.992-7.424-65.728-11.52a256.128 256.128 0 0 0-346.496-346.432a1736.64 1736.64 0 0 1-11.392-65.728z"/>`),
		g.Group(children),
	)
}