package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExcelLogoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 8c0-1.105.806-2 1.8-2h23.4c.994 0 1.8.895 1.8 2v32c0 1.105-.806 2-1.8 2H16.8c-.994 0-1.8-.895-1.8-2v-6H7.8c-.994 0-1.8-.895-1.8-2V16c0-1.105.806-2 1.8-2H15V8Zm13 6V8H17v6h11Zm2-6v6h10V8H30Zm-2 8h-4v7h4v-7Zm2 7v-7h10v7H30Zm-2 2h-4v7h4v-7Zm2 7v-7h10v7H30Zm-2 2H17v6h11v-6Zm2 6v-6h10v6H30Zm-8-24v16H7.8V16H22Zm-10.774 3h2.147l1.743 3.754L16.957 19h2.006L16.2 24l2.827 5H16.91l-1.899-3.93l-1.89 3.93h-2.148l2.874-5.018L11.226 19Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}