package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm121.3 316.5c8.4 8.4 8.4 22 0 30.4l-30.4 30.4c-8.4 8.4-22 8.4-30.4 0l-60.8-60.8l-60.8 60.8c-8.4 8.4-22 8.4-30.4 0L134 346.9c-8.4-8.4-8.4-22 0-30.4l60.8-60.8l-60.8-60.9c-8.4-8.4-8.4-22 0-30.4l30.4-30.4c8.4-8.4 22-8.4 30.4 0l60.8 60.8l60.8-60.8c8.4-8.4 22-8.4 30.4 0l30.4 30.4c8.4 8.4 8.4 22 0 30.4l-60.8 60.8l60.9 60.9z"/>`),
		g.Group(children),
	)
}