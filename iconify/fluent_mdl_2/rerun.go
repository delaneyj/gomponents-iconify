package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rerun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 640V512h421l-107-107l-108-108q-81-80-182-122t-216-43q-79 0-151 21t-136 58t-116 91t-89 117t-58 137t-21 151q0 112 42 214t122 181l599 599v182l-690-690q-97-97-149-224t-53-265q0-96 25-185t72-167t110-142t143-109t167-71T963 5q125 0 219 34t174 91t153 133t155 158V0h128v640h-640zm128 512l768 448l-768 448v-896zm128 223v450l386-225l-386-225z"/>`),
		g.Group(children),
	)
}