package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InputComposite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 27v85h43v128H0V112h43V27q0-9 6-15.5T64 5t15 6.5T85 27zm86 298v-42h128v42q0 21-12 37.5T256 385v90h-43v-90q-19-6-30.5-22.5T171 325zM0 325v-42h128v42q0 21-12 37.5T85 385v90H43v-90q-19-6-31-22.5T0 325zm427-213h42v128H341V112h43V27q0-9 6.5-15.5t15-6.5t15 6.5T427 27v85zM256 27v85h43v128H171V112h42V27q0-9 6.5-15.5t15-6.5t15 6.5T256 27zm85 298v-42h128v42q0 21-11.5 37.5T427 385v90h-43v-90q-19-6-31-22.5T341 325z"/>`),
		g.Group(children),
	)
}