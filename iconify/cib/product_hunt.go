package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductHunt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.394 13.675A2.329 2.329 0 0 1 18.069 16h-4.394v-4.65h4.394a2.329 2.329 0 0 1 2.325 2.325zM31.5 16c0 8.563-6.938 15.5-15.5 15.5S.5 24.562.5 16C.5 7.437 7.438.5 16 .5S31.5 7.438 31.5 16zm-8.006-2.325a5.428 5.428 0 0 0-5.425-5.425h-7.494v15.5h3.1V19.1h4.394a5.428 5.428 0 0 0 5.425-5.425z"/>`),
		g.Group(children),
	)
}