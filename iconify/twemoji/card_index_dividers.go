package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIndexDividers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#FDCB58" d="M31 15a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h22a2 2 0 0 1 2 2v8z"/><path fill="#FDD888" d="M33 19a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h26a2 2 0 0 1 2 2v8z"/><path fill="#FEE7B8" d="M35 33a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2V15a2 2 0 0 1 2-2h30a2 2 0 0 1 2 2v18z"/><path fill="#67757F" d="M11 14zm0 0z"/><path fill="#78B159" d="M23 3h-5a2 2 0 0 0-2 2v1h2a2 2 0 0 1 2 2h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"/><path fill="#F18F26" d="M17 7h-5a2 2 0 0 0-2 2v1h2a2 2 0 0 1 2 2h3a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/><path fill="#9268CA" d="M13 14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h5a2 2 0 0 1 2 2v1z"/>`),
		g.Group(children),
	)
}