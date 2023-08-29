package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lounge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.39 24.35H23a1 1 0 0 1 1 1v5h0H8.38h0v-5a1 1 0 0 1 1.01-1Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.38 25.36V15.28a1.41 1.41 0 0 1 1.41-1.41h12.8A1.41 1.41 0 0 1 24 15.28v10.08m0 0V15.28a1.41 1.41 0 0 1 1.41-1.41h12.8a1.41 1.41 0 0 1 1.41 1.41v10.08M5.36 19.57h2.16a.86.86 0 0 1 .86.86v10h0h-2.81a1.07 1.07 0 0 1-1.07-1.1v-8.9a.86.86 0 0 1 .86-.86Zm37.07 10.84h-2.81h0v-10a.86.86 0 0 1 .86-.86h2.16a.86.86 0 0 1 .86.86v8.9a1.07 1.07 0 0 1-1.07 1.07Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25 24.35h13.61a1 1 0 0 1 1 1v5h0H24h0v-5a1 1 0 0 1 1-1ZM8.38 30.41h3.07v3.73H8.38zm28.17 0h3.07v3.73h-3.07z"/>`),
		g.Group(children),
	)
}