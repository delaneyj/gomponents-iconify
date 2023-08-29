package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#93a2aa" d="M36.6 15.1c-1.7 2.6-6.1 5.3-10 7.3c.2.2.6.8.8.8c1.1-.5 2.2-.9 3.3-1.4c2.2-1 4.4-2.3 6.4-3.8c4-3.2 5.5-12.8 8.4-13c1.7 0 3.3 2.5 4.5 3.6l1.7-1.7c.1-.1-1.3-1.4-1.4-1.5c-1.7-1.7-3.7-3.5-6.4-2.5c-3.6 1.4-5.3 9.3-7.3 12.2m9.3-10s.1 0 0 0m-1.3.2c-.1 0-.1 0 0 0M2.3 60.9c-.3.4-.4.8-.2 1c.2.2.7.1 1-.2L5.5 60L4 58.5l-1.7 2.4"/><g fill="#42ade2"><path d="m28.3 26.3l9.4 9.4l23.4-25l-7.8-7.8z"/><path d="m9.372 46.795l19.729-19.728l7.778 7.778L17.15 54.573z"/></g><path fill="#c7d3d8" d="m3.2 57.7l3.1 3.1l10.9-6.2l-7.8-7.8zm28-34.3l9.4 9.4l-2.9 2.9l-9.4-9.4z"/><path fill="#42ade2" d="m53.8 3.4l6.8 6.8c1.9-1.9 1.9-4.9 0-6.8s-4.9-1.9-6.8 0"/><path fill="#c7d3d8" d="m51.953 3.748l1.414-1.414l8.344 8.344l-1.414 1.414z"/>`),
		g.Group(children),
	)
}