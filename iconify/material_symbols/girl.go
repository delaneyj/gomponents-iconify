package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Girl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7.5q-.725 0-1.238-.513T10.25 5.75q0-.725.513-1.238T12 4q.725 0 1.238.513t.512 1.237q0 .725-.513 1.238T12 7.5ZM10 20v-4H8l2.375-6.375q.2-.5.638-.813T12 8.5q.55 0 .988.313t.637.812L16 16h-2v4h-4Z"/>`),
		g.Group(children),
	)
}