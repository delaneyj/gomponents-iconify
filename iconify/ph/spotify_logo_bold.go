package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpotifyLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 20a108 108 0 1 0 108 108A108.12 108.12 0 0 0 128 20Zm0 192a84 84 0 1 1 84-84a84.09 84.09 0 0 1-84 84Zm66.59-84.36a12 12 0 0 1-16.24 4.93a106.7 106.7 0 0 0-100.7 0A11.83 11.83 0 0 1 72 134a12 12 0 0 1-5.66-22.58a130.61 130.61 0 0 1 123.3 0a12 12 0 0 1 4.95 16.22Zm-16 36a12 12 0 0 1-16.23 5a73 73 0 0 0-68.72 0a12 12 0 0 1-11.28-21.18a97 97 0 0 1 91.28 0a12 12 0 0 1 4.95 16.18Z"/>`),
		g.Group(children),
	)
}