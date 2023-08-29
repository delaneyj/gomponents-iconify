package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overtime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#CFD8DC" d="M12 40V20h32v20c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4z"/><path fill="#78909C" d="M44 16v6H12v-6c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4z"/><g fill="#37474F"><circle cx="37" cy="16" r="3"/><circle cx="20" cy="16" r="3"/></g><path fill="#B0BEC5" d="M37 10c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2s2-.9 2-2v-4c0-1.1-.9-2-2-2zm-17 0c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2s2-.9 2-2v-4c0-1.1-.9-2-2-2z"/><path fill="#90A4AE" d="M32 34h4v4h-4zm-6 0h4v4h-4zm-6 0h4v4h-4zm12-6h4v4h-4zm-6 0h4v4h-4zm-6 0h4v4h-4z"/><circle cx="16" cy="15" r="12" fill="#F44336"/><circle cx="16" cy="15" r="9" fill="#eee"/><path d="M15 8h2v7h-2z"/><path d="m20.518 18.1l-1.343 1.344l-3.818-3.818l1.344-1.343z"/><circle cx="16" cy="15" r="1.5"/>`),
		g.Group(children),
	)
}