package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Performance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22.007H5v-2h10ZM22 4l-4.735 5.955l-.22.27l-5.63 7.19a2.001 2.001 0 1 1-2.83-2.83ZM2.645 7.234A10.843 10.843 0 0 0 1.19 15H2v-1a9.685 9.685 0 0 1 1.69-5.52ZM12 2a10.958 10.958 0 0 0-8.119 3.596L5.025 6.96A7.428 7.428 0 0 1 10 5a7.432 7.432 0 0 1 4.997 1.978l3.55-2.802A10.936 10.936 0 0 0 12 2Zm6.83 9.2l-.233.287l-.728.93A10.118 10.118 0 0 1 18 14v1h4.81a10.879 10.879 0 0 0-1.183-7.318Z"/>`),
		g.Group(children),
	)
}