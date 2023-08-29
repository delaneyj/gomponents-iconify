package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Debitum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.58 33.295a6.61 6.61 0 0 0 1.878 1.692m10.258-10.302A5.927 5.927 0 0 0 25.951 24h-3.903a5.954 5.954 0 0 1-5.942-5.954h0a5.954 5.954 0 0 1 5.943-5.954h3.528a6.615 6.615 0 0 1 5.843 2.613M24 12.092V9.115"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.72 41.677L41.694 11.89M27.308 37.782h6.615m-6.615-4.746v1.99h1.99l8.532-8.533l-1.99-1.99Zm10.522-6.543l2.08-2.08l-1.99-1.99l-2.08 2.08"/>`),
		g.Group(children),
	)
}