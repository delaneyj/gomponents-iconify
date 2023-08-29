package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOrEdgeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 17.175V11H3q-.825 0-1.413-.588T1 9V4q0-.425.288-.713T2 3q.425 0 .713.288T3 4v5h3.5q.825 0 1.413.588T8.5 11v6.175l.875-.875q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712L8.2 20.3q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.275-.275-.288-.688T4.2 16.3q.275-.275.7-.275t.7.275l.9.875Zm11 .025l.875-.9q.275-.3.688-.288t.712.288q.3.3.3.713t-.3.712L17.2 20.3q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.275-.275-.287-.688t.287-.712q.275-.275.7-.275t.7.275l.9.875V11q0-.825.588-1.413T17.5 9H21V4q0-.425.288-.713T22 3q.425 0 .713.288T23 4v5q0 .825-.588 1.413T21 11h-3.5v6.2Z"/>`),
		g.Group(children),
	)
}