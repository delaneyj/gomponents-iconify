package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixWallpapers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.001 16.552l-11.12 19.26c-.94 1.63.235 3.666 2.116 3.666h22.24c1.882 0 3.058-2.036 2.117-3.666l-11.12-19.26c-.94-1.63-3.292-1.63-4.233 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.223 30.024L15.53 25.36c-.723-1.25-2.529-1.25-3.251 0L5.754 36.663c-.722 1.251.18 2.815 1.626 2.815h9.617M6.19 18.996h6.023m25.05-1.222H42.5"/><circle cx="16.49" cy="12.057" r="3.535" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}