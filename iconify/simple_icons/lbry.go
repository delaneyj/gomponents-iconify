package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lbry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23.35 14.19l.168-1.066l-1.031-.177l.067-.414l1.446.245l-.237 1.48zm.151-5.496v1.192l-11.734 7.211l-8.842-4.336l.017-.668l8.792 4.328L22.91 9.557v-.49l-10.55-5.09l-11.768 7.28v3.254l11.142 5.512l11.632-7.135l.33.507L11.767 20.7L0 14.883v-3.956L12.325 3.3z"/>`),
		g.Group(children),
	)
}