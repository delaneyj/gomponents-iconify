package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallenLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#ffd93b" d="M56.5 31.3c0-.5.1-.9.3-1.3c-1.9-1.1-3-2.7-3-4.5c.1-2.8 3.3-5.1 7.4-5.5c1-9.4.8-17 .8-17S24.9 1.7 11.4 14.7l-.7.7c2 1.2 3.3 3.2 3.3 5.5c0 3.7-3.5 6.8-7.8 6.8c-1.3 0-2.5-.3-3.6-.8c-3.9 11.7 12 15.1 15.5 18.5c4.5 4.3 9 28.4 31.7 6.5c4.1-4 6.9-10.2 8.7-16.8c-1.2-1-2-2.3-2-3.8"/><path fill="#db9523" d="m24.5 40.9l4.6-4.6l19 11.7l-18.3-12.5l7.6-7.6l17 7.5l-16.3-8.3l7.1-7.2l12.9-.6l-11.9-.4L62 3L45.5 18.2l-.4-11.5l-.6 12.5l-7.4 6.9l-8.6-15.8l7.8 16.5l-8 7.3l-12.9-17.6l12.2 18.3l-4.8 4.4l-12.9-7.6l12.2 8.3L9.6 51.4l2.3 2.2l11.9-12l8.6 11.7z"/>`),
		g.Group(children),
	)
}