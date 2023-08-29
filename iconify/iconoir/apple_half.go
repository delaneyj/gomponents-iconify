package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m12.147 21.265l-.147-.03l-.147.03c-2.377.475-4.62.21-6.26-1.1C3.964 18.86 2.75 16.373 2.75 12c0-4.473 1.008-6.29 2.335-6.954c.695-.347 1.593-.448 2.735-.317c1.141.132 2.458.488 3.943.983l.26.086l.255-.102c2.482-.992 4.713-1.373 6.28-.641c1.47.685 2.692 2.538 2.692 6.945c0 4.374-1.213 6.86-2.843 8.164c-1.64 1.312-3.883 1.576-6.26 1.1Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 5.5C12 3 11 2 9 2"/><path d="M12 6v15"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 12v2"/></g>`),
		g.Group(children),
	)
}