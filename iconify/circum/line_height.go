package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.439 4.062h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1Zm0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1Zm0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1ZM3.208 18.8a.5.5 0 0 1 .71-.71l1.14 1.14V4.775l-1.14 1.14a.513.513 0 0 1-.71 0a.5.5 0 0 1 0-.71l2-2a.494.494 0 0 1 .34-.14a.549.549 0 0 1 .37.14l2 2a.524.524 0 0 1 0 .71a.5.5 0 0 1-.71 0l-1.15-1.15v14.47l1.15-1.15a.5.5 0 1 1 .71.71l-2 2a.513.513 0 0 1-.71 0Z"/>`),
		g.Group(children),
	)
}