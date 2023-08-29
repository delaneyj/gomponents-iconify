package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nucleus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M544 320L736 0l160 128l-288 224zM128 128L288 0l192 320l-64 32zm288 352L64 709L0 512l384-96zm160 32l64 512H384l64-512h128zm64-96l384 96l-64 197l-352-229z"/>`),
		g.Group(children),
	)
}