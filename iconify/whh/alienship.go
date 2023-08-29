package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alienship(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 640q-139 0-257-21.5T68.5 560T0 479.5t69.5-80T256 342v-86q0-106 75-181T512 0t181 75t75 181v86q117 21 186.5 57.5t69.5 80t-68.5 80.5T769 618.5T512 640zm192-367q0-86-56-147.5T512 64t-136 61.5T320 273v68q90 43 192 43t192-43v-68zm254 349q-15 62-143 104t-303 42t-303-42T66 622q67 38 186 60t260 22t260-22t186-60z"/>`),
		g.Group(children),
	)
}