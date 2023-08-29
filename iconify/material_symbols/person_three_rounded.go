package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonThreeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14q-1.25 0-2.125-.875T7 11q0-.55.175-1.038t.525-.887q-.1-.25-.15-.525T7.5 8q0-.95.513-1.688T9.35 5.226q.5-.575 1.175-.9T12 4q.8 0 1.475.325t1.175.9q.825.35 1.338 1.088T16.5 8q0 .275-.05.55t-.15.525q.35.4.525.888T17 11q0 1.25-.875 2.125T14 14h-4Zm-4 8q-.825 0-1.413-.588T4 20v-.8q0-.85.438-1.563T5.6 16.55q1.55-.775 3.15-1.163T12 15q1.65 0 3.25.388t3.15 1.162q.725.375 1.163 1.088T20 19.2v.8q0 .825-.588 1.413T18 22H6Z"/>`),
		g.Group(children),
	)
}