package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.45 15.5q.6.6 1.55.588t1.4-.688L19 7l-8.4 5.6q-.675.45-.712 1.375t.562 1.525ZM12 4q1.475 0 2.838.413T17.4 5.65l-1.9 1.2q-.825-.425-1.712-.637T12 6Q8.675 6 6.337 8.337T4 14q0 1.05.288 2.075T5.1 18h13.8q.575-.95.838-1.975T20 13.9q0-.9-.213-1.75t-.637-1.65l1.2-1.9q.75 1.175 1.188 2.5T22 13.85q.025 1.425-.325 2.725T20.65 19.05q-.275.45-.75.7t-1 .25H5.1q-.525 0-1-.25t-.75-.7q-.65-1.125-1-2.388T2 14q0-2.075.788-3.888t2.15-3.174Q6.3 5.575 8.124 4.787T12 4Zm.175 7.825Z"/>`),
		g.Group(children),
	)
}