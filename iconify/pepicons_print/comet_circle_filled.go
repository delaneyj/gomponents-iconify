package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CometCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M26 14c0 6.627-5.373 12-12 12S2 20.627 2 14S7.373 2 14 2s12 5.373 12 12Z" opacity=".2"/><g fill-rule="evenodd" clip-rule="evenodd"><path d="M8.66 16.418a1.5 1.5 0 1 0 2.28-1.95a1.5 1.5 0 0 0-2.28 1.95Zm2.765.925a2.5 2.5 0 1 1-3.25-3.8a2.5 2.5 0 0 1 3.25 3.8Z"/><path d="M13.327 19.573a5.421 5.421 0 1 1-7.029-8.253l7.08-5.836c.719-.593 1.786.033 1.62.95l-.51 2.8l5.617-4.68c.87-.725 2.084.324 1.494 1.289l-3.39 5.547l1.91-.212c.97-.108 1.504 1.1.77 1.745l-7.562 6.65Zm-6.94-1.196a4.42 4.42 0 0 0 6.28.445l7.562-6.65l-1.91.212a1 1 0 0 1-.964-1.515l3.39-5.548l-5.616 4.681c-.718.598-1.792-.028-1.624-.947l.509-2.8l-7.08 5.837a4.42 4.42 0 0 0-.546 6.285Z"/></g><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}