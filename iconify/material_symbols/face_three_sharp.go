package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceThreeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q3.35 0 5.675-2.337T20 10q0-.775-.125-1.488t-.4-1.387q-.675.425-1.425.65T16.5 8q-1.35 0-2.538-.613T12 5.65q-.775 1.125-1.963 1.738T7.5 8q-.8 0-1.55-.225t-1.425-.65q-.275.675-.4 1.388T4 10q0 3.325 2.337 5.663T12 18Zm-3-5.75q.525 0 .888-.363T10.25 11q0-.525-.363-.888T9 9.75q-.525 0-.888.363T7.75 11q0 .525.363.888T9 12.25Zm6 0q.525 0 .888-.363T16.25 11q0-.525-.363-.888T15 9.75q-.525 0-.888.363T13.75 11q0 .525.363.888t.887.362ZM0 22L1.1 9.95q.2-2.1 1.137-3.925t2.4-3.163Q6.1 1.525 7.988.762T12 0q2.125 0 4.013.763t3.35 2.1q1.462 1.337 2.4 3.162T22.9 9.95L24 22H0Z"/>`),
		g.Group(children),
	)
}