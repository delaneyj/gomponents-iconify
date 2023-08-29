package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Traingo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.456 7.879h21.088a2.478 2.478 0 0 1 2.483 2.484V34.57a2.478 2.478 0 0 1-2.483 2.484H13.456a2.478 2.478 0 0 1-2.483-2.484V10.363a2.478 2.478 0 0 1 2.482-2.484Zm2.615 32.401H31.93M14.27 43.5l3.604-6.447m12.25 0l3.604 6.447"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.01 7.879h19.979v15.464H14.01z"/><circle cx="17.629" cy="28.462" r="3.246" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="30.371" cy="28.462" r="3.246" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.521 4.5H31.48m-4.118 0v3.379M20.638 4.5v3.379m-9.665 25.638h26.054m-5.159-4.787a1.515 1.515 0 0 0-2.994 0"/><circle cx="16.762" cy="27.954" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.068 33.517c-.039.992-.008 1.482-.682 1.897h-8.772c-.674-.415-.643-.905-.682-1.897"/>`),
		g.Group(children),
	)
}