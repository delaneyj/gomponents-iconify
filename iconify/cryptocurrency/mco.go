package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mco(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.02-26.982h-.066L6.5 10.501v11l9.414 5.48l.077.019l9.509-5.5V10.501l-9.52-5.483zm-.031 1.138l1.031.572l7.52 4.35v9.845l-7.52 4.325l-1.032.597l-8.448-4.92v-9.849l8.449-4.92zm-7.14 10.61l3.41 5.96h1.362l1.612-1.51v-.756l-1.672-1.612v-2.54l-2.21-1.413l-2.502 1.872zm7.903 4.452l1.61 1.491h1.344l3.393-5.942l-2.496-1.889l-2.187 1.43v2.54l-1.667 1.61l.003.76zm-2.37-4.91l-.25 2.39l1.845-.004l1.867-.004l-.236-2.382l.798-2.131h-4.838l.813 2.131zm1.595-2.715l4.622.003l-.901-3.825h-7.464l-.855 3.82l4.598.002z"/>`),
		g.Group(children),
	)
}