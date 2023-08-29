package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zfont(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.16 30.219h29.679V42.5H9.16zm29.68-5.117H9.16l-3.07 5.117h35.82l-3.07-5.117zm-7.669-4.108v1.038H16.829L31.171 5.5H16.829v1.023"/>`),
		g.Group(children),
	)
}