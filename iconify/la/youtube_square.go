package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm9 4.08s-4.39 0-5.47.29c-.6.17-1.08.64-1.24 1.25C9 13.71 9 16.01 9 16.01s0 2.3.29 3.4c.16.6.64 1.06 1.24 1.22c1.08.29 5.47.29 5.47.29s4.39 0 5.47-.29c.6-.17 1.08-.62 1.24-1.23c.29-1.09.29-3.39.29-3.39s0-2.29-.29-3.39c-.16-.61-.64-1.08-1.24-1.25c-1.08-.29-5.47-.29-5.47-.29zm-2 2.076l5 2.856l-5 2.842v-5.698z"/>`),
		g.Group(children),
	)
}