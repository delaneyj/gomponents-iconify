package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossedFlags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fff" d="m6.84 25.478l19.85-9.489l6.47 13.533l-19.85 9.489zM58.528 39.01L38.68 29.523l6.47-13.533l19.848 9.489z"/><circle cx="20" cy="27.5" r="2.5"/><circle cx="51.841" cy="27.5" r="2.5"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M46 56L27 16M6.84 25.478l19.85-9.489l6.47 13.533l-19.85 9.489zM25.841 56l19-40m13.687 23.01L38.68 29.523l6.47-13.533l19.848 9.489z"/><g fill="#d22f27"><circle cx="20" cy="27.5" r="2.5"/><circle cx="51.841" cy="27.5" r="2.5"/></g>`),
		g.Group(children),
	)
}