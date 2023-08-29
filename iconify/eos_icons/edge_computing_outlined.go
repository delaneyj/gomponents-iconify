package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeComputingOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.09 13a1.5 1.5 0 1 0 0-1h-2.8A5.448 5.448 0 0 0 14 10.37V5.91a1.5 1.5 0 1 0-1 0v4.18a5.551 5.551 0 0 0-1-.09a5.61 5.61 0 0 0-2.28.48L8 8.76V6h1V2H6v4h1v3.17l1.82 1.82a5.562 5.562 0 0 0-1.81 2.04a4.418 4.418 0 0 0-1.03.24l-2.13-2.13a1.519 1.519 0 1 0-.71.71l1.89 1.89A4.499 4.499 0 0 0 7.5 22h9.75a3.74 3.74 0 0 0 .26-7.47a5.358 5.358 0 0 0-.54-1.53Zm-1.72 3.53A1.736 1.736 0 0 1 19 18.25A1.758 1.758 0 0 1 17.25 20H7.5a2.498 2.498 0 0 1-.28-4.98l1.07-.11l.5-.96a3.618 3.618 0 0 1 6.76.97l.3 1.5Z"/>`),
		g.Group(children),
	)
}