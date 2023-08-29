package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderHeart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.75 23.676a4.176 4.176 0 0 0-4.176-4.176A4.169 4.169 0 0 0 24 21.525a4.17 4.17 0 0 0-3.574-2.025a4.176 4.176 0 0 0-4.176 4.176c0 .527.102 1.03.28 1.494C17.91 29.142 24 33.5 24 33.5s6.09-4.358 7.47-8.33c.179-.465.28-.967.28-1.494h0Z"/>`),
		g.Group(children),
	)
}