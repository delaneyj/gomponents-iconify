package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingButWallpapers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45 23.723c0 11.598-9.402 21-21 21s-21-9.403-21-21h42Zm-6.183 0l-6.9-8.417l-6.901 8.417"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.586 20.587L19.521 8.822L9.306 23.723"/><circle cx="34.192" cy="6.719" r="3.441" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}