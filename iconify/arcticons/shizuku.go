package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shizuku(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 42.5h-33a2 2 0 0 1-2-2v-33a2 2 0 0 1 2-2h33a2 2 0 0 1 2 2v33a2 2 0 0 1-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.31 33.96h1.44L35.5 24l-5.75-9.96h-11.5l-2.81 4.87"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.65 42.5l1.92-5.28a13.31 13.31 0 0 0 .16-8.63l.34-.62a12.13 12.13 0 0 0 1.14-6.35a12 12 0 0 0-5.92.3l-.21.07a13.5 13.5 0 0 0-3.49-1.9h0a13.53 13.53 0 0 0-3.91-.79l-.11-.19a12.13 12.13 0 0 0-4.34-4A12 12 0 0 0 8 20.67q-.09.34-.15.69a13.71 13.71 0 0 0-2.35 1.92"/><circle cx="12.35" cy="25.53" r="1.1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="21.63" cy="28.91" r="1.1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}