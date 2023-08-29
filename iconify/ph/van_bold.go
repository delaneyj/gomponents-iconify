package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VanBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m253.11 112.18l-45.57-53.09A19.94 19.94 0 0 0 192.26 52H32a20 20 0 0 0-20 20v112a20 20 0 0 0 20 20h14.06a36 36 0 0 0 67.88 0h40.12a36 36 0 0 0 67.88 0H236a20 20 0 0 0 20-20v-64a12.05 12.05 0 0 0-2.89-7.82ZM217.89 108H176V76h14.42ZM104 108V76h48v32ZM80 76v32H36V76Zm0 128a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm108 0a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm33.94-24a36 36 0 0 0-67.88 0h-40.12a36 36 0 0 0-67.88 0H36v-48h196v48Z"/>`),
		g.Group(children),
	)
}