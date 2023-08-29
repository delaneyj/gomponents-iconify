package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TMobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M18.297 1.536h.771c4.989 0 7.312 2.609 8.281 9.568l1.547-.068L28.688 0H3.312l-.249 11.036l1.484.068c.26-2.609.568-4.016 1.24-5.416c1.183-2.609 3.651-4.152 6.692-4.152h1.079v24.016c0 2.541-.151 3.344-.771 3.948c-.516.469-1.541.667-2.729.667H8.875v1.604h14.104v-1.604h-1.187c-1.177 0-2.213-.197-2.724-.667c-.615-.604-.771-1.407-.771-3.948V1.536zM2.667 14.5h6.505v6.495H2.667zm19.912 0h6.504v6.495h-6.504z"/>`),
		g.Group(children),
	)
}