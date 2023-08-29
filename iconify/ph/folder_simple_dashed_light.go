package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderSimpleDashedLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M124.4 84.8L94.53 62.4a2 2 0 0 0-1.2-.4H40a2 2 0 0 0-2 2v16a6 6 0 0 1-12 0V64a14 14 0 0 1 14-14h53.33a14 14 0 0 1 8.4 2.8l29.87 22.4a6 6 0 0 1-7.2 9.6ZM88 202H39.38a1.4 1.4 0 0 1-1.38-1.38V192a6 6 0 0 0-12 0v8.62A13.39 13.39 0 0 0 39.38 214H88a6 6 0 0 0 0-12Zm72 0h-32a6 6 0 0 0 0 12h32a6 6 0 0 0 0-12Zm64-56a6 6 0 0 0-6 6v48.89a1.11 1.11 0 0 1-1.11 1.11H200a6 6 0 0 0 0 12h16.89A13.12 13.12 0 0 0 230 200.89V152a6 6 0 0 0-6-6Zm-8-72h-48a6 6 0 0 0 0 12h48a2 2 0 0 1 2 2v24a6 6 0 0 0 12 0V88a14 14 0 0 0-14-14ZM32 158a6 6 0 0 0 6-6v-32a6 6 0 0 0-12 0v32a6 6 0 0 0 6 6Z"/>`),
		g.Group(children),
	)
}