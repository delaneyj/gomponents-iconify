package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fsharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M0 64.5L60.7 3.8v30.4L30.4 64.5l30.4 30.4v30.4L0 64.5zm39.1 0l21.7-21.7v43.4L39.1 64.5zm88.9 0L65.1 3.8v30.4l30.4 30.4l-30.4 30.3v30.4L128 64.5z"/>`),
		g.Group(children),
	)
}