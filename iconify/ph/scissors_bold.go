package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M238.78 183.79L98.28 87.65A40.18 40.18 0 0 0 100 76a40 40 0 1 0-15.29 31.45l30 20.56l-30 20.56a40 40 0 1 0 3.57 59.74A39.73 39.73 0 0 0 100 180a40.18 40.18 0 0 0-1.72-11.66l37.72-25.8l89.22 61.06a12 12 0 0 0 13.56-19.81Zm-167.47 7.54A16 16 0 1 1 76 180a16 16 0 0 1-4.69 11.33ZM48.69 87.3a16 16 0 1 1 22.62 0a16 16 0 0 1-22.62 0Zm112.82 23.24a12 12 0 0 1 3.13-16.68l60.58-41.46a12 12 0 0 1 13.56 19.81l-60.59 41.46a12 12 0 0 1-16.68-3.13Z"/>`),
		g.Group(children),
	)
}