package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hangouts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.997 0a6.91 6.91 0 0 0-6.909 6.909c0 3.616 3.294 6.547 6.909 6.547V16c4.197-2.128 6.916-5.556 6.916-9.091c0-3.816-3.1-6.909-6.916-6.909zM7 8c0 .828-.447 1.5-1 1.5V8H4V5h3v3zm5 0c0 .828-.447 1.5-1 1.5V8H9V5h3v3z"/>`),
		g.Group(children),
	)
}