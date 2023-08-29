package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M0 16a7.418 7.418 0 0 0 7.422 7.417a7.418 7.418 0 1 0 0-14.834C3.328 8.583 0 11.901 0 16zm17.156 0a7.418 7.418 0 0 0 7.422 7.417C28.661 23.417 32 20.099 32 16s-3.323-7.417-7.417-7.417c-4.104 0-7.427 3.318-7.427 7.417z"/>`),
		g.Group(children),
	)
}