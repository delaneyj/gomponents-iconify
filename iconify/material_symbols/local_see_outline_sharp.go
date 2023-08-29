package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalSeeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.5q1.35 0 2.438-.713t1.637-1.862q-.4-.35-.8-.738t-.8-.812q-.15.95-.837 1.538T12 16.5q-1.05 0-1.775-.725T9.5 14q0-1.05.738-1.763t1.787-.737q-.3-.45-.575-.913t-.525-.937q-1.475.375-2.45 1.575T7.5 14q0 1.875 1.313 3.188T12 18.5ZM9 4h1.25q-.125.475-.188.975T10 6h-.125L8.05 8H4v12h16v-4.125q.5-.45 1-.938t1-.987V22H2V6h5.15L9 4Zm.5 10H12H9.5Zm8.5 1q3.025-2.575 4.513-4.775T24 6.15q0-2.825-1.813-4.487T18 0q-2.375 0-4.188 1.663T12 6.15q0 1.875 1.488 4.075T18 15Zm0-2.675q-2.6-2.45-3.3-3.95T14 6.15q0-2.025 1.263-3.088T18 2q1.475 0 2.738 1.063T22 6.15q0 .725-.7 2.225t-3.3 3.95ZM16.15 9l.7-2.275L15 5.25h2.275L18 3l.725 2.25H21l-1.85 1.475l.7 2.275L18 7.6L16.15 9ZM18 7.15Z"/>`),
		g.Group(children),
	)
}