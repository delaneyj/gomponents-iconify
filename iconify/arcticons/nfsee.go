package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nfsee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="12.833" cy="24.019" r="3.74" fill="none" stroke="currentColor"/><path fill="none" stroke="currentColor" d="M20.726 35.567a3.693 3.693 0 0 1-2.617-6.297a7.397 7.397 0 0 0 2.168-5.251c0-1.972-.773-3.84-2.177-5.26a3.692 3.692 0 0 1 5.252-5.191c2.78 2.81 4.309 6.522 4.309 10.45c0 3.948-1.533 7.663-4.318 10.461a3.683 3.683 0 0 1-2.617 1.088Z"/><path fill="none" stroke="currentColor" d="M28.697 43.5a3.693 3.693 0 0 1-2.617-6.297a18.568 18.568 0 0 0 5.443-13.184c0-4.993-1.944-9.687-5.474-13.216A3.692 3.692 0 0 1 31.27 5.58c4.925 4.925 7.637 11.473 7.637 18.438c0 6.941-2.697 13.473-7.593 18.394a3.685 3.685 0 0 1-2.617 1.087Z"/>`),
		g.Group(children),
	)
}