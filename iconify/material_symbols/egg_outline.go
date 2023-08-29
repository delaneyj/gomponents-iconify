package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.925 0-4.963-2.038T5 14q0-1.925.638-3.875t1.65-3.538Q8.3 5 9.55 4T12 3q1.225 0 2.463 1t2.25 2.588q1.012 1.587 1.65 3.537T19 14q0 2.925-2.038 4.963T12 21Zm0-2q2.075 0 3.538-1.463T17 14q0-1.425-.488-3t-1.225-2.913q-.737-1.337-1.612-2.212T12 5q-.775 0-1.663.875T8.713 8.087Q7.975 9.425 7.488 11T7 14q0 2.075 1.463 3.538T12 19Zm1-1q.425 0 .713-.288T14 17q0-.425-.288-.713T13 16q-1.25 0-2.125-.875T10 13q0-.425-.288-.713T9 12q-.425 0-.713.288T8 13q0 2.075 1.463 3.538T13 18Zm-1-6Z"/>`),
		g.Group(children),
	)
}