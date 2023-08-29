package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.534 1.866a8.502 8.502 0 0 1 11.487 9.793l5.154 5.153l-1.415 1.414l-5.987-5.987l.178-.576a6.498 6.498 0 0 0-1.615-6.518A6.496 6.496 0 0 0 8.34 3.392l4.228 4.228l-4.95 4.95L3.39 8.34a6.496 6.496 0 0 0 1.754 5.996a6.498 6.498 0 0 0 6.518 1.615l.575-.178l5.988 5.988l-1.414 1.414l-5.154-5.153A8.502 8.502 0 0 1 1.864 6.535l.254-.624l1.672.002l3.828 3.828L9.74 7.62L5.912 3.792L5.91 2.12l.624-.254Zm9.216 12.471l4.95 4.95l-1.414 1.414l-4.95-4.95l1.414-1.414Z"/>`),
		g.Group(children),
	)
}