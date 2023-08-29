package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addcomment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 896q-66 0-134-16q-34 40-69.5 69.5t-60 43.5t-47.5 21.5t-30.5 8.5t-10.5 1q26-57 30-124.5T177 786Q94 723 47 635T0 448q0-91 40.5-174t109-143T313 35.5T512 0t199 35.5T874.5 131t109 143t40.5 174t-40.5 174t-109 143T711 860.5T512 896zm256-480q0-13-9.5-22.5T736 384H576V224q0-13-9.5-22.5T544 192h-64q-13 0-22.5 9.5T448 224v160H288q-13 0-22.5 9.5T256 416v64q0 13 9.5 22.5T288 512h160v160q0 13 9.5 22.5T480 704h64q13 0 22.5-9.5T576 672V512h160q13 0 22.5-9.5T768 480v-64z"/>`),
		g.Group(children),
	)
}