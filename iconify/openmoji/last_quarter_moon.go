package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M64 36A28 28 0 0 0 36 8v56a28 28 0 0 0 28-28Z"/><path fill="#fcea2b" d="M21.549 30.124a1.5 1.5 0 0 0 0-3a1.5 1.5 0 0 0 0 3Z"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" d="M8 36a28 28 0 0 0 28 28V8A28 28 0 0 0 8 36Z"/><path fill="#3f3f3f" stroke="#3f3f3f" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M64 36A28 28 0 0 0 36 8v56a28 28 0 0 0 28-28Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><path d="M8 36a28 28 0 0 0 28 28V8A28 28 0 0 0 8 36Z"/></g>`),
		g.Group(children),
	)
}