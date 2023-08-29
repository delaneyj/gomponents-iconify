package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNuxt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.146 8.583l-1.3-2.09a1.046 1.046 0 0 0-1.786.017l-5.91 9.908A1.046 1.046 0 0 0 4.047 18H7.96m12.083 0c.743 0 1.201-.843.82-1.505l-4.044-7.013a.936.936 0 0 0-1.638 0l-4.043 7.013c-.382.662.076 1.505.819 1.505h8.086z"/>`),
		g.Group(children),
	)
}