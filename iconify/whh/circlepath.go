package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circlepath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-768q-50 0-94.5 13T336 307.5t-58.5 70T256 480q0 65 40 120l43-43q-7-11-13-38t-6-39q0-67 60.5-113.5T512 320q68 0 130 40.5t62 87.5q0 51-34.5 89.5T576 576q-19 0-27.5-.5t-19-3.5t-14-9.5T512 544v-96q0-13-9.5-22.5T480 416q-13 0-22.5 9.5T448 448v224q0 11-4.5 18.5t-15 10t-18 3t-22.5.5h-4q-14 0-22 8t-9 16l-1 8q0 12 3.5 19t13 9.5t16.5 3t23 .5h40q26 0 45-18.5t19-45.5v-96q0 12 21.5 22t42.5 10q92 0 142-54t50-138q0-61-38-106t-95-65.5T512 256z"/>`),
		g.Group(children),
	)
}