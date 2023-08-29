package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorriedFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30c16.571 0 30.002-13.432 30.002-30S48.57 2 32 2zm0 57.5C16.836 59.5 4.5 47.164 4.5 32S16.836 4.5 32 4.5c15.165 0 27.502 12.336 27.502 27.5S47.164 59.5 32 59.5z"/><circle cx="20.5" cy="29.59" r="5" fill="currentColor"/><circle cx="43.502" cy="29.59" r="5" fill="currentColor"/><path fill="currentColor" d="M52.619 17.051a16.418 16.418 0 0 0-13.492-3.615c-.703.135-.193 2.27.385 2.156c4.17-.748 8.457.4 11.693 3.133c.443.386 1.955-1.205 1.414-1.674m-28.131-1.61c.578.113 1.09-2.021.387-2.156A16.425 16.425 0 0 0 11.383 16.9c-.541.469.969 2.063 1.412 1.674a14.236 14.236 0 0 1 11.693-3.133m16.094 28.987c-5.404-2.538-11.788-2.54-17.197-.012c-1.339.645.329 4.15 1.662 3.5c3.571-1.665 8.896-2.306 13.875.01c1.334.619 3.078-2.813 1.66-3.498"/>`),
		g.Group(children),
	)
}