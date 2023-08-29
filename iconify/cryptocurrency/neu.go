package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm0-28C9.375 4 4 9.375 4 16s5.375 12 12 12s12-5.375 12-12S22.625 4 16 4zm-.05 19.62l-6.476-3.84v-7.668l6.477-3.83l6.476 3.83v7.669l-6.476 3.838zm-2.925-12.872l-.132.082L19 20.48v-9.658l-3.05-1.808l-2.925 1.734zm-.198 10.282l3.124 1.858l2.852-1.693l-5.976-9.444v9.28zM10.1 12.482v6.937l2.178 1.29v-9.517l-2.178 1.29zm9.452 8.269l2.252-1.332v-6.937l-2.252-1.331v9.6z"/>`),
		g.Group(children),
	)
}