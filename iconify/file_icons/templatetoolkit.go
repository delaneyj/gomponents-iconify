package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Templatetoolkit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m83.62 46.544l353.47-2.646c110.008 10.36 92.868 178.97-12.368 180.224l-77.172 1.847v163.045c.093 101.333-174.376 107.79-177.199 4.86v-168.88H93.652c-117.969 6.632-128-171.818-10.031-178.45zm210.758 342.47V170.943h129.764c40.668 0 49.123-68.785 10.216-73.904l-349.4 2.667c-46.389 3.805-40.584 73.153 5.921 71.237h128.174v222.633c6.987 30.68 77.326 28.48 75.325-4.562z"/>`),
		g.Group(children),
	)
}