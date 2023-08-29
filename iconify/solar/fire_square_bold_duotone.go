package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireSquareBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M3.464 20.535C4.93 22 7.286 22 12 22c4.714 0 7.071 0 8.535-1.465C22 19.072 22 16.714 22 12s0-7.071-1.465-8.536C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464C2 4.93 2 7.286 2 12c0 4.714 0 7.071 1.464 8.535Z" opacity=".5"/><path d="M17 12.667C17 16.933 13.444 18 11.667 18C10.11 18 7 16.933 7 12.667C7 10.81 8.063 9.633 8.956 9.04c.408-.271.916-.098.942.391c.058 1.071.883 1.931 1.523 1.07c.584-.788.873-1.858.873-2.501c0-.947.958-1.548 1.707-.968C15.459 8.162 17 10.056 17 12.667Z"/></g>`),
		g.Group(children),
	)
}