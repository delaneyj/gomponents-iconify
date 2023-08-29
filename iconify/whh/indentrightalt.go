package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Indentrightalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M904 640H568q-50 0-85-37.5T448 512t35-90.5t85-37.5h336q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5zm0-384H248q-50 0-85-37.5T128 128t35-90.5T248 0h656q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5zm-712 71q7-8 15.5-8t15.5 8l155 167q6 8 6 18t-6 17L223 697q-7 7-15.5 7t-15.5-7V576H32q-13 0-22.5-9.5T0 544v-64q0-13 9.5-22.5T32 448h160V327zm56 441h656q50 0 85 37.5t35 90.5t-35 90.5t-85 37.5H248q-50 0-85-37.5T128 896t35-90.5t85-37.5z"/>`),
		g.Group(children),
	)
}