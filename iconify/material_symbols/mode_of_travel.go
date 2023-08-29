package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModeOfTravel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-4.025-3.425-6.012-6.362T4 10.2q0-3.75 2.413-5.975T12 2q3.175 0 5.588 2.225T20 10.2l2.1-2.1l1.4 1.4L19 14l-4.5-4.5l1.4-1.4l2.1 2.1q0-2.725-1.738-4.462T12 4Q9.475 4 7.737 5.738T6 10.2q0 1.775 1.475 4.063T12 19.35q.5-.45.925-.875l.85-.85q-.125-.25-.2-.537T13.5 16.5q0-1.05.725-1.775T16 14q1.05 0 1.775.725T18.5 16.5q0 1.05-.725 1.775T16 19q-.2 0-.363-.025T15.3 18.9q-.725.75-1.538 1.525T12 22Z"/>`),
		g.Group(children),
	)
}