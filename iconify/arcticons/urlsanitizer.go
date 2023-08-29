package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Urlsanitizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.055 14.211l6.942-6.942c2.679-2.678 7.539-2.162 10.217.517s3.195 7.538.517 10.217L30.448 28.286c-2.679 2.678-7.538 2.162-10.217-.517"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.945 33.788l-6.942 6.943c-2.679 2.678-7.539 2.162-10.217-.517s-3.195-7.538-.517-10.217l10.283-10.283c2.679-2.678 7.538-2.162 10.217.517"/>`),
		g.Group(children),
	)
}