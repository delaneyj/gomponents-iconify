package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gasbuddy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.415 30.073c2.518-.562 13.264-2.31 14.413-7.09c.845-3.512-12.964-6.636-12.09-9.88m13.95-4.01c-2.034.484-6.035 1.14-5.754 2.5c.558 2.693 12.396 2.74 18.191 7.09c4.476 3.36 8.553 11.91-3.429 20.224"/><path fill="none" stroke="currentColor" stroke-dasharray="4 4" stroke-linecap="round" stroke-linejoin="round" d="M21.712 35.71c5.821-3.594 9.828-8.65 9.357-11.623c-1.122-7.092-19.72-7.815-17.412-12.671"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}