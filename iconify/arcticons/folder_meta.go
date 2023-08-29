package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderMeta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 11.5a3 3 0 0 1 3-3h8.718a4 4 0 0 1 2.325.745l4.914 3.51a4 4 0 0 0 2.325.745H40.5a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3h-33a3 3 0 0 1-3-3v-25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 24.308c-1.025-1.651-2.268-2.914-3.701-2.914S16 23.292 16 27.734c0 3.274.883 3.872 1.822 3.872c2.42 0 4.992-5.002 6.178-7.298Zm0 0c1.025-1.651 2.268-2.914 3.701-2.914S32 23.292 32 27.734c0 3.274-.883 3.872-1.822 3.872c-2.42 0-4.992-5.002-6.178-7.298Z"/>`),
		g.Group(children),
	)
}