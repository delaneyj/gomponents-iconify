package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLocationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 9a3.75 3.75 0 1 1 7.5 0a3.75 3.75 0 0 1-7.5 0ZM12 6.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M5.456 8.127a6.403 6.403 0 0 1 6.382-5.877h.324a6.403 6.403 0 0 1 6.382 5.877a6.895 6.895 0 0 1-1.534 4.931l-3.595 4.396a1.828 1.828 0 0 1-2.83 0L6.99 13.058a6.895 6.895 0 0 1-1.534-4.93Zm6.382-4.377a4.903 4.903 0 0 0-4.887 4.5a5.395 5.395 0 0 0 1.2 3.859l3.595 4.396c.131.16.377.16.508 0l3.595-4.396a5.395 5.395 0 0 0 1.2-3.859a4.903 4.903 0 0 0-4.887-4.5h-.324Z" clip-rule="evenodd"/><path fill="currentColor" d="M7.67 16.335a.75.75 0 1 0-1.34-.67l-2 4A.75.75 0 0 0 5 20.75h14a.75.75 0 0 0 .67-1.085l-2-4a.75.75 0 1 0-1.34.67l1.457 2.915H6.213l1.457-2.915Z"/>`),
		g.Group(children),
	)
}