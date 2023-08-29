package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 0c8.837 0 16 7.163 16 16s-7.163 16-16 16S0 24.837 0 16S7.163 0 16 0zM7.5 7.257l7.595 19.112l10.583-12.894l-3.746-3.562L7.5 7.257zm16.252 6.977l-7.67 9.344l.983-8.133l6.687-1.21zM9.472 9.488l6.633 5.502l-1.038 8.58L9.472 9.487zM21.7 11.083l2.208 2.099l-6.038 1.093l3.83-3.192zM10.194 8.778l10.402 1.914l-4.038 3.364l-6.364-5.278z"/>`),
		g.Group(children),
	)
}