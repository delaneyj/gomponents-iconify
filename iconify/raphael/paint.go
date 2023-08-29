package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25.54 5.772V2.208H3.293v8.083h22.25V6.77c1.21.114 2.166 1.175 2.166 2.48v1.374c0 1.398-1.164 2.784-2.54 3.025l-7.883 1.38c-1.857.326-3.37 2.125-3.37 4.01v.386c-.74.366-1.29 1.507-1.29 2.865v4.5c0 1.65.806 3 1.79 3s1.793-1.35 1.793-3v-4.5c0-1.357-.55-2.498-1.292-2.864v-.385c0-1.397 1.164-2.783 2.54-3.024l7.883-1.38c1.856-.327 3.368-2.126 3.368-4.01V9.25c0-1.856-1.404-3.364-3.167-3.478z"/>`),
		g.Group(children),
	)
}