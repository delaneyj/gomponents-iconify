package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phorum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 895v43q0 37-31 64t-74 21q-55-7-110-26q-51-19-78-46t-27-75V761q-61-62-94.5-142.5T0 448q0-91 35.5-174T131 131t143-95.5T448 0t174 35.5T765 131t95 143t35 174t-35 173.5T765 764t-143 95.5T448 895zm0-767q-69 0-165 26.5T102 226q-20 11-11 85q8 67 26 73q4 1 7-1q9-3 32-15.5t36-15.5v511q0 19 19 45t44 35q33 12 66 16q26 4 44.5-12.5T384 908V757q55 10 96 10q87 0 168.5-42T781 608t51-160q0-67-24.5-124.5T736 222t-121-69t-167-25zm22 448q-29 0-86-8V326q46-6 70-6q89 0 137.5 36t48.5 92q0 52-52.5 90T470 576z"/>`),
		g.Group(children),
	)
}