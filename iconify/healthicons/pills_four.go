package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillsFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M22 14.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-3.39-.325a2 2 0 1 0-2.209-3.334l-6.012 3.984a2 2 0 1 0 2.21 3.335l6.012-3.985ZM41 14.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-6.874-4.076a2 2 0 1 0-3.752 1.387l2.5 6.765a2 2 0 1 0 3.752-1.387l-2.5-6.765ZM22 33.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-7.956 4.099a2 2 0 0 0 3.691-1.541l-2.779-6.656a2 2 0 1 0-3.691 1.54l2.779 6.657ZM41 33.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-3.463-.84a2 2 0 0 0-2.612-3.03l-5.462 4.71a2 2 0 0 0 2.612 3.03l5.462-4.71Z"/>`),
		g.Group(children),
	)
}