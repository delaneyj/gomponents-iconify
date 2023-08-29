package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bubble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.495 21.495 0 0 0 5.096 13.75h37.808A21.495 21.495 0 0 0 24 2.5Zm0 43a21.486 21.486 0 0 0 19.041-11.534H4.96A21.486 21.486 0 0 0 24 45.5Zm18.904-31.75H5.096a21.48 21.48 0 0 0-.137 20.216H43.04a21.48 21.48 0 0 0-.137-20.216Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.084 24a21.404 21.404 0 0 0-2.596-10.25m.137 20.216A21.388 21.388 0 0 0 38.084 24m-28.168-.284a21.404 21.404 0 0 0 2.596 10.25m-.137-20.216a21.388 21.388 0 0 0-2.459 9.966"/><ellipse cx="24" cy="17.225" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.36" ry="3.475"/>`),
		g.Group(children),
	)
}