package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 22.744a3.756 3.756 0 1 1 0 7.512a3.756 3.756 0 0 1 0-7.512Zm0-3.244v.919m7 6.081h-.919M24 33.5v-.919M17 26.5h.919m11.031-4.95l-.65.65m.65 9.25l-.65-.65m-9.25.65l.65-.65m-.65-9.25l.65.65"/>`),
		g.Group(children),
	)
}