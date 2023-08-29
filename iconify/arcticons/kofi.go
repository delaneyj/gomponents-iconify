package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kofi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 9.853h28.295v19.225a9.07 9.07 0 0 1-9.07 9.07H13.57a9.07 9.07 0 0 1-9.07-9.07V9.853h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.197 25.95l6.45 6.45l6.45-6.45a8.2 8.2 0 0 0 2.402-5.799h0a4.508 4.508 0 0 0-4.212-4.546a4.426 4.426 0 0 0-4.64 4.421a4.426 4.426 0 0 0-4.64-4.42a4.508 4.508 0 0 0-4.212 4.545h0a8.2 8.2 0 0 0 2.402 5.798ZM32.795 9.853h2.935a7.77 7.77 0 0 1 0 15.54h-2.935"/>`),
		g.Group(children),
	)
}