package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5.66 13.418a1.5 1.5 0 1 0 2.28-1.95a1.5 1.5 0 0 0-2.28 1.95Zm2.765.925a2.5 2.5 0 1 1-3.25-3.8a2.5 2.5 0 0 1 3.25 3.8Z"/><path d="M10.327 16.573A5.421 5.421 0 1 1 3.298 8.32l7.08-5.836c.719-.593 1.786.033 1.62.95l-.51 2.8l5.617-4.68c.87-.725 2.084.324 1.494 1.289l-3.39 5.547l1.91-.212c.97-.108 1.504 1.1.77 1.745l-7.562 6.65Zm-6.94-1.196a4.42 4.42 0 0 0 6.28.445l7.562-6.65l-1.91.212a1 1 0 0 1-.964-1.515l3.39-5.548l-5.616 4.681c-.718.598-1.792-.028-1.624-.947l.509-2.8l-7.08 5.837a4.42 4.42 0 0 0-.546 6.285Z"/></g>`),
		g.Group(children),
	)
}