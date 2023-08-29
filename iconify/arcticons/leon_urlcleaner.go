package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeonUrlcleaner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.485 23.718a.8.8 0 0 0-1.008.777c.03 2.951 1.264 8.812 5.462 12.719c.699.65 8.165-3.486 8.165-3.486l-4.265 6.538a18.12 18.12 0 0 0 4.606 2.219a.388.388 0 0 0 .394-.11c.918-.98 5.45-5.976 6.534-10.39c0 0-7.633-7.403-9.193-9.732c0 0-3.537 2.07-6.293 2.07a18.362 18.362 0 0 1-4.402-.605Zm10.695-1.465c1.924-1.088 5.443-3.593 7.072-2.175a26.494 26.494 0 0 1 3.516 3.46c1.11 1.45-.423 5.133-1.395 8.447"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.984 21.626C29.388 18.57 42.523 5.5 42.523 5.5"/>`),
		g.Group(children),
	)
}