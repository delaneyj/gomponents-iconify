package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Banuser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M800 1024q-93 0-158.5-65.5T576 800t65.5-158.5T800 576t158.5 65.5T1024 800t-65.5 158.5T800 1024zm160-224q0-50-29-91L709 931q41 29 91 29q66 0 113-47t47-113zm-320 0q0 50 29 91l222-222q-41-29-91-29q-66 0-113 47t-47 113zm-128 0q0 66 28.5 124.5T619 1024H436l-81-1l-87-2.5l-80-4.5l-75.5-8.5l-57.5-12L13.5 979L0 957q2-88 110-155.5T384 712v-33q-52-23-90-65t-60-98.5t-32-121T192 256q0-65 25-114.5t69-80t101-46T512 0t125 15.5t101 46t69 80T832 256q0 156-37 256q-118 2-200.5 86T512 800z"/>`),
		g.Group(children),
	)
}