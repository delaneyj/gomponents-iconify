package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405 85h-42V43q0-18-13-30.5T320 0H43Q25 0 12.5 12.5T0 43v256q0 17 12.5 29.5T43 341h4q6 19 22.5 31t37.5 12q20 0 36.5-12t22.5-31h180q6 19 22.5 31t36.5 12q21 0 37.5-12t22.5-31h4q18 0 30.5-12.5T512 299v-77q0-17-13-30zM107 341q-8 0-15-6.5T85 320t7-14.5t15-6.5t14.5 6.5T128 320t-6.5 14.5T107 341zm213-42H166q-6-19-22.5-31T107 256q-21 0-37.5 12T47 299h-4v-86h277v86zm0-214v86H43V43h277v42zm85 256q-8 0-14.5-6.5T384 320t6.5-14.5T405 299t15 6.5t7 14.5t-7 14.5t-15 6.5zm64-42h-4q-6-19-22.5-31T405 256q-27 0-42 17V128h12l94 94v77z"/>`),
		g.Group(children),
	)
}