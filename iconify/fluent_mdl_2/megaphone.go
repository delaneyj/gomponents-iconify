package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 768q14 0 38 1t48 1h20q9 0 16-2l1798-320q6-1 16-1t22-1q24 0 50 1t40 1v1152h-39q-25 0-49 1h-30q-13 0-21-2l-763-136q-10 57-38 106t-71 84t-94 55t-111 20q-66 0-124-25t-102-68t-69-102t-25-125q0-28 6-57l-396-71q-8-1-17-1t-21-1q-22 0-44 1t-40 1V768zm832 832q35 0 67-12t57-33t42-51t23-64l-378-67q-3 18-3 35q0 40 15 75t41 61t61 41t75 15zM1920 577L128 896v256q7 0 11 1l1781 318V577z"/>`),
		g.Group(children),
	)
}