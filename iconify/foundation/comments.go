package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comments(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M94.974 17.838a2.526 2.526 0 0 0-2.443-2.277v-.007H39.624a2.53 2.53 0 0 0-2.513 2.284h-.026v13.139h22.342c1.174 0 2.127.896 2.243 2.039h.023v29.232h11.009l6.488 6.487a2.537 2.537 0 0 0 4.644-1.415v-5.072h8.698v-.008a2.534 2.534 0 0 0 2.417-2.027H95V17.838h-.026z"/><path fill="currentColor" d="M54.44 35.964H7.204v.006a2.256 2.256 0 0 0-2.181 2.033H5v37.835h.046a2.264 2.264 0 0 0 2.158 1.809v.006h7.765v4.529a2.266 2.266 0 0 0 4.146 1.263l5.792-5.792H54.44c1.096 0 2.011-.78 2.22-1.815h.045V38.003h-.023a2.255 2.255 0 0 0-2.242-2.039z"/>`),
		g.Group(children),
	)
}