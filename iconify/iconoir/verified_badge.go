package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerifiedBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.528 1.6a.6.6 0 0 1 .944 0l1.809 2.3a.6.6 0 0 0 .635.207l2.815-.798a.6.6 0 0 1 .764.554l.11 2.925a.6.6 0 0 0 .393.54l2.747 1.01a.6.6 0 0 1 .292.897l-1.63 2.431a.6.6 0 0 0 0 .668l1.63 2.431a.6.6 0 0 1-.292.897l-2.747 1.01a.6.6 0 0 0-.392.54l-.111 2.925a.6.6 0 0 1-.764.554l-2.815-.798a.6.6 0 0 0-.636.206L12.473 22.4a.6.6 0 0 1-.944 0L9.72 20.1a.6.6 0 0 0-.635-.207l-2.815.798a.6.6 0 0 1-.764-.554l-.11-2.925a.6.6 0 0 0-.393-.54l-2.747-1.01a.6.6 0 0 1-.292-.897l1.63-2.431a.6.6 0 0 0 0-.668l-1.63-2.431a.6.6 0 0 1 .292-.897l2.747-1.01a.6.6 0 0 0 .392-.54l.111-2.925a.6.6 0 0 1 .764-.554l2.815.798A.6.6 0 0 0 9.72 3.9l1.81-2.3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 12l2 2l4-4"/></g>`),
		g.Group(children),
	)
}