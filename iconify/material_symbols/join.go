package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Join(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19q-.7 0-1.363-.138t-1.287-.387q1.475-1.35 2.313-3.013T16.5 12q0-1.8-.838-3.463T13.35 5.525q.625-.25 1.288-.388T16 5q2.925 0 4.963 2.038T23 12q0 2.925-2.038 4.963T16 19Zm-4-1.25q-1.375-.95-2.188-2.45T9 12q0-1.8.813-3.3T12 6.25q1.375.95 2.188 2.45T15 12q0 1.8-.813 3.3T12 17.75ZM8 19q-2.925 0-4.963-2.038T1 12q0-2.925 2.038-4.963T8 5q.7 0 1.363.138t1.287.387q-1.475 1.35-2.313 3.013T7.5 12q0 1.975.813 3.663t2.237 2.862q-.6.225-1.238.35T8 19Z"/>`),
		g.Group(children),
	)
}