package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.357 11.011A354.291 354.291 0 0 0 4.5 36.297m17.275-24.792a420.801 420.801 0 0 0-9.264 25.088"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.995 25.912a126.727 126.727 0 0 1 16.153-3.495m3.033 14.176a146.067 146.067 0 0 1 8.67-25.087a221.918 221.918 0 0 0-1.285 22.714A163.931 163.931 0 0 1 43.5 11.505a228.682 228.682 0 0 0-4.648 25.484m-16.5-5.983c-1.352 2.093-2.92 2.619-3.577 1.78c-.478-.61.372-1.89.824-2.456c.891-1.116 2.292-1.703 2.538-2.786c.215-.94-.461-1.104-1.104-1.104c-.775 0-1.256.966-1.137 1.417c.15 1.38 1.601 4.93 2.027 5.753"/>`),
		g.Group(children),
	)
}