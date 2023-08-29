package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlFoodBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 100h-4.78a92 92 0 0 0-182.44 0H32a12 12 0 0 0-12 12a108.38 108.38 0 0 0 56 94.68V208a20 20 0 0 0 20 20h64a20 20 0 0 0 20-20v-1.32A108.38 108.38 0 0 0 236 112a12 12 0 0 0-12-12Zm-53.71-39.94a92 92 0 0 0-43.1 39.94H106a68.27 68.27 0 0 1 62-40c.76 0 1.52 0 2.29.06Zm17.22 19.08a67.66 67.66 0 0 1 7.41 20.86h-38.79a67.91 67.91 0 0 1 31.38-20.86ZM128 44c.83 0 1.65 0 2.48.06A92.3 92.3 0 0 0 80.37 100H61.08A68.1 68.1 0 0 1 128 44Zm35 144.39a12 12 0 0 0-7 10.91v4.7h-56v-4.7a12 12 0 0 0-7-10.91A84.32 84.32 0 0 1 44.87 124h166.26A84.32 84.32 0 0 1 163 188.39Z"/>`),
		g.Group(children),
	)
}