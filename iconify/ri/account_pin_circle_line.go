package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountPinCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.745 21.745C5.308 20.722 2 16.747 2 12C2 6.477 6.477 2 12 2s10 4.477 10 10c0 4.747-3.308 8.722-7.745 9.745L12 24l-2.255-2.255Zm-2.733-3.488a7.953 7.953 0 0 0 3.182 1.539l.56.129L12 21.172l1.247-1.247l.56-.13a7.955 7.955 0 0 0 3.36-1.686A6.979 6.979 0 0 0 12.16 16c-2.036 0-3.87.87-5.148 2.257ZM5.616 16.82A8.975 8.975 0 0 1 12.16 14a8.972 8.972 0 0 1 6.362 2.634a8 8 0 1 0-12.906.187ZM12 13a4 4 0 1 1 0-8a4 4 0 0 1 0 8Zm0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}