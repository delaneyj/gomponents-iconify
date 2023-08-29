package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m214.662 100.752l-75.97 316.401H0v91.456h47.573v590.639h136.569v-96.685h831.716v96.685h136.496V508.609H1200v-91.456h-138.692l-75.971-316.401H214.662zm40.474 61.52h689.713l61.187 254.887H193.934l61.186-254.887h.016zm-168.048 423.7l190.44 70.329v91.384l-190.44-70.329v-91.384zm1025.73 0v91.384l-190.362 70.329v-91.456l190.362-70.257zM415.833 825.036h368.193v91.384H415.833v-91.384z"/>`),
		g.Group(children),
	)
}