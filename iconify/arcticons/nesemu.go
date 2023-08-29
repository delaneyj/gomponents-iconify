package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nesemu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.206 29.495v3.89c0 1.297 1.062 2.358 2.358 2.358s2.358-1.06 2.358-2.358v-3.89m0 4.008v2.358m-16.638-4.008c0-1.297 1.061-2.358 2.358-2.358S24 30.555 24 31.853v3.89m-4.716-6.248v6.366M24 31.853c0-1.297 1.061-2.358 2.358-2.358s2.358 1.06 2.358 2.358v3.89m-16.638-4.598h3.065m1.651 4.716h-4.716V26.43h4.716m-5.842-4.164V12.139l6.71 10.127V12.139m2.743 5.063h3.291m1.773 5.064h-5.064V12.139h5.064m2.87 8.988c.633.76 1.392 1.139 2.532 1.139h1.519c1.392 0 2.531-1.14 2.531-2.532s-1.139-2.532-2.531-2.532h-1.646c-1.393 0-2.532-1.139-2.532-2.532s1.14-2.531 2.532-2.531h1.519c1.14 0 1.899.253 2.532 1.139m2.753 8.988a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/>`),
		g.Group(children),
	)
}