package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Support(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896q-66 0-134-16q-34 40-69.5 69.5t-60 43.5t-47.5 21.5t-30 8.5t-11 1q26-57 30-124.5T176 786Q94 723 47 635T0 448q0-91 40.5-174t109-143T313 35.5T512 0t199 35.5T874.5 131t109 143t40.5 174t-40.5 174t-109 143T711 860.5T512 896zm-64-160q0 13 9 22.5t23 9.5h64q13 0 22.5-9.5T576 736v-64q0-14-9.5-23t-22.5-9h-64q-14 0-23 9t-9 23v64zm64-608q-85 0-152 37.5T268 262l116 58q0-27 37.5-45.5T512 256t90.5 18.5t37.5 45t-37.5 45.5t-90.5 19q-27 0-45.5 18.5T448 448v96q0 13 9 22.5t23 9.5h64q13 0 22.5-9.5T576 544v-39q83-16 137.5-67.5T768 320q0-80-75-136t-181-56z"/>`),
		g.Group(children),
	)
}