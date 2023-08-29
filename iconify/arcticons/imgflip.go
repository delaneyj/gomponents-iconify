package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imgflip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.12 24.215a5.73 5.73 0 0 1 5.747-5.71h0a5.73 5.73 0 0 1 5.747 5.71v9.138M15.119 18.505v14.848m11.495-9.138a5.73 5.73 0 0 1 5.747-5.71h0a5.73 5.73 0 0 1 5.748 5.71v9.138"/><ellipse cx="11.814" cy="15.472" fill="currentColor" rx=".755" ry=".75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.814 18.505v14.848"/>`),
		g.Group(children),
	)
}