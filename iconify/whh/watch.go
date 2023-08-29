package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1020.82 855l-165 165q-12 12-23 0l-219-218l-25 25q-6 6-14.5 6t-14.5-6l-58-58q-99-5-168.5-75t-74.5-168l-59-59q-6-6-6-14.5t6-14.5l25-25l-218-219q-12-11 0-23l165-165q12-12 23 0l219 218l25-25q6-6 14.5-6t14.5 6l59 59q98 5 168 74.5t75 168.5l58 58q6 6 6 14.5t-6 14.5l-25 25l218 219q12 11 0 23zm-506.5-534q-79.5 0-136 56.5t-56.5 136t56.5 135.5t136 56t135.5-56t56-135.5t-56-136t-135.5-56.5zm22.5 278q-9 10-23 10t-24-10l-94-93q-10-10-10-23.5t9.5-23.5t23.5-10t24 10l71 70l71-70q9-10 23-10t24 10t10 23.5t-10 23.5z"/>`),
		g.Group(children),
	)
}