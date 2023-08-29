package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonFourOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q-1.65 0-2.825-1.175T8 9V5.5q0-.625.438-1.063T9.5 4q.375 0 .713.175t.537.5q.2-.325.537-.5T12 4q.375 0 .713.175t.537.5q.2-.325.537-.5T14.5 4q.625 0 1.063.438T16 5.5V9q0 1.65-1.175 2.825T12 13Zm0-2q.825 0 1.413-.588T14 9V6.5h-4V9q0 .825.588 1.413T12 11ZM6 21q-.825 0-1.413-.588T4 19v-.8q0-.85.438-1.563T5.6 15.55q1.55-.775 3.15-1.163T12 14q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T20 18.2v.8q0 .825-.588 1.413T18 21H6Zm0-2h12v-.8q0-.275-.138-.5t-.362-.35q-1.35-.675-2.725-1.012T12 16q-1.4 0-2.775.338T6.5 17.35q-.225.125-.363.35T6 18.2v.8Zm6 0Zm0-12.5Z"/>`),
		g.Group(children),
	)
}