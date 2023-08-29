package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VanThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m251 117.4l-45.53-53.1a12 12 0 0 0-9.21-4.3H32a12 12 0 0 0-12 12v112a12 12 0 0 0 12 12h20.29a28 28 0 0 0 55.42 0h56.58a28 28 0 0 0 55.42 0H240a12 12 0 0 0 12-12v-64a4 4 0 0 0-1-2.6Zm-51.64-47.93L239.3 116H172V68h24.26a4 4 0 0 1 3.1 1.47ZM100 116V68h64v48ZM32 68h60v48H28V72a4 4 0 0 1 4-4Zm48 144a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm112 0a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm48-24h-20.29a28 28 0 0 0-55.42 0h-56.58a28 28 0 0 0-55.42 0H32a4 4 0 0 1-4-4v-60h216v60a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}