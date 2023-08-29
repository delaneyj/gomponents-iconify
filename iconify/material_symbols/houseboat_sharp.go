package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseboatSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-2q.95 0 1.4-.5t1.925-.5q1.45 0 1.95.5t1.375.5q.95 0 1.413-.5T12 16q1.475 0 1.938.5t1.412.5q.875 0 1.375-.5t1.95-.5q1.475 0 1.925.5t1.4.5v2q-1.425 0-1.95-.5t-1.4-.5q-.9 0-1.4.5t-1.925.5q-1.475 0-1.925-.5T12 18q-.95 0-1.4.5t-1.925.5q-1.425 0-1.925-.5t-1.4-.5q-.875 0-1.4.5T2 19Zm2.5-4l-2.25-2.25l1.4-1.4L5.325 13H7V9.625l-1.325.975L4.5 9L12 3.5L19.5 9l-1.175 1.625L17 9.65V13h1.675l1.675-1.65l1.4 1.4L19.5 15h-15Zm6.5-2h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}