package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bearish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#F44336" d="M40 34h4v10h-4zm-6-5h4v15h-4zm-6 4h4v11h-4zm-6-8h4v19h-4zm-6 3h4v16h-4zm-6-4h4v20h-4zm-6-5h4v25H4z"/><g fill="#D32F2F"><path d="m34 13.2l-4 4l-10-10l-5 5l-7.6-7.6l-2.8 2.8L15 17.8l5-5l10 10l4-4l6.1 6.1l2.8-2.8z"/><path d="M44 26h-9l9-9z"/></g>`),
		g.Group(children),
	)
}