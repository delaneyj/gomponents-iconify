package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssistantDirectionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23q-2.275 0-4.288-.863t-3.5-2.35q-1.487-1.487-2.35-3.5T1 12q0-2.3.863-4.3t2.35-3.488q1.487-1.487 3.5-2.35T12 1q2.3 0 4.3.863t3.488 2.35Q21.275 5.7 22.138 7.7T23 12q0 2.275-.863 4.288t-2.35 3.5q-1.487 1.487-3.487 2.35T12 23Zm0-11Zm-.05 8.35l8.375-8.375L11.95 3.6l-8.375 8.375l8.375 8.375Zm-3.975-5.375v-5h5.15l-1.05-1.1l1.4-1.4l3.5 3.5l-3.5 3.5l-1.4-1.4l1.05-1.1h-3.15v3h-2ZM12 21q3.775 0 6.388-2.613T21 12q0-3.775-2.613-6.388T12 3Q8.225 3 5.612 5.613T3 12q0 3.775 2.613 6.388T12 21Z"/>`),
		g.Group(children),
	)
}