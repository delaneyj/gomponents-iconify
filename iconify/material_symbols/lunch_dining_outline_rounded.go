package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LunchDiningOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10q-.425 0-.713-.288T2 9q0-2.875 2.713-4.438T12 3q4.575 0 7.288 1.563T22 9q0 .425-.288.713T21 10H3Zm1.15-2h15.7q-.575-1.45-2.663-2.225T12 5q-3.1 0-5.188.775T4.15 8ZM2 13.5q0-.35.25-.688t.7-.562q.45-.225.95-.488t1.45-.262q1.4 0 1.9.5t1.4.5q.9 0 1.425-.5T12 11.5q1.4 0 1.925.5t1.425.5q.9 0 1.4-.5t1.9-.5q.925 0 1.45.25t.95.475q.425.225.688.563T22 13.5q0 .425-.363.675t-.737.15q-.75-.2-1.088-.512T18.75 13.5q-.9 0-1.45.5t-1.95.5q-1.4 0-1.925-.5T12 13.5q-.9 0-1.425.5t-1.925.5q-1.4 0-1.9-.5t-1.4-.5q-.725 0-1.1.313t-1.15.512q-.375.1-.738-.15T2 13.5ZM4 21q-.825 0-1.413-.588T2 19v-1q0-.825.588-1.413T4 16h16q.825 0 1.413.588T22 18v1q0 .825-.588 1.413T20 21H4Zm0-2h16v-1H4v1Zm.15-11h15.7h-15.7ZM4 18h16H4Z"/>`),
		g.Group(children),
	)
}