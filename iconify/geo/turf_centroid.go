package geo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TurfCentroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<circle cx="50.178" cy="49.646" r="7.796" fill="currentColor"/><path fill="currentColor" d="M85.342 21.482a2 2 0 0 1-2-2v-3h-3a2 2 0 0 1 0-4h5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2zm-13.618-5h-8.618a2 2 0 0 1 0-4h8.618a2 2 0 0 1 0 4zm-17.237 0H45.87a2 2 0 0 1 0-4h8.618a2 2 0 0 1-.001 4zm-17.235 0h-8.618a2 2 0 0 1 0-4h8.618a2 2 0 0 1 0 4zm-22.236 5a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h5a2 2 0 0 1 0 4h-3v3a2 2 0 0 1-2 2zm0 51.708a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2zm0-17.236a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2zm0-17.235a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2zm5 48.09h-5a2 2 0 0 1-2-2v-5a2 2 0 0 1 4 0v3h3a2 2 0 0 1 0 4zm51.707 0h-8.617a2 2 0 0 1 0-4h8.617a2 2 0 0 1 0 4zm-17.236 0h-8.618a2 2 0 0 1 0-4h8.618a2 2 0 0 1 0 4zm-17.236 0h-8.618a2 2 0 0 1 0-4h8.618a2 2 0 0 1 0 4zm48.091 0h-5a2 2 0 0 1 0-4h3v-3a2 2 0 0 1 4 0v5a2 2 0 0 1-2 2zm0-13.619a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2zm0-17.236a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2zm0-17.235a2 2 0 0 1-2-2v-8.618a2 2 0 0 1 4 0v8.618a2 2 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}