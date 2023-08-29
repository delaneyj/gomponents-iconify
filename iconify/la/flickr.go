package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5C6.355 5 5 6.355 5 8v16c0 1.645 1.355 3 3 3h16c1.645 0 3-1.355 3-3V8c0-1.645-1.355-3-3-3zm0 2h16c.566 0 1 .434 1 1v16c0 .566-.434 1-1 1H8c-.566 0-1-.434-1-1V8c0-.566.434-1 1-1zm3.5 5.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 1 0 0-7zm9 0a3.5 3.5 0 1 0 0 7a3.5 3.5 0 1 0 0-7z"/>`),
		g.Group(children),
	)
}