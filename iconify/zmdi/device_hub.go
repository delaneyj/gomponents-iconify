package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceHub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 277h85v107H277v-65l-85-90l-85 90v65H0V277h85l86-85v-68q-19-7-31-23.5T128 64q0-27 18.5-45.5T192 0t45.5 18.5T256 64q0 20-12 36.5T213 124v68z"/>`),
		g.Group(children),
	)
}