package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PregnantWomen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M33 4v7.519c7.312 4.194 10.614 9.775 9.906 16.742C42.198 35.228 38.896 40.474 33 44M13.772 4c-2.138.37-3.677 1.622-4.616 3.758C7.746 10.96 5 21.963 5 23.726c0 1.762 3.568 6.509 12.052 13.71c3.931 3.336 6.255 3.166 7.513.928s-.407-3.674-2.515-5.327c-3.863-3.029-8.948-7.822-8.948-9.926c0-1.403 1.317-5.677 3.95-12.824"/><path d="M8.201 28.94a95.28 95.28 0 0 0 2.8 15M33 19.944c1.547 1.028 2.547 2.365 3 4.012c.453 1.648.585 3.15.396 4.506"/></g>`),
		g.Group(children),
	)
}