package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.942 13.5v7.148M19.982 5.5h3.573m-3.573 3.574h2.319M19.982 5.5v7.148m4.004 4V9.5h2.338a2.4 2.4 0 0 1 0 4.801h-2.338m2.338 0l2.337 2.347m4.502.852h4.73m-2.365 7.148V17.5m2.222 4h4.729l-4.729 7.148h4.729"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.75 30.961a58.51 58.51 0 0 1-5.947-1.048a2.936 2.936 0 0 0-2.832.548c-.5.51-1.945 1.935-3.446 3.408a27.087 27.087 0 0 1-12.368-12.377c1.463-1.502 2.832-2.937 3.39-3.437a2.936 2.936 0 0 0 .547-2.832a58.378 58.378 0 0 1-1.048-5.958a1.888 1.888 0 0 0-2.073-1.684q-.05.005-.099.013H6.83a1.416 1.416 0 0 0-1.284 1.294c-.519 7.26 3.531 15.03 4.353 16.522v.056l.113.217a33.46 33.46 0 0 0 12.273 12.273l.416.237c1.888 1 9.394 4.777 16.409 4.258a1.416 1.416 0 0 0 1.312-1.285v-8.034a1.888 1.888 0 0 0-1.572-2.158q-.05-.008-.099-.013ZM28.09 5.5A30 30 0 0 1 42.5 17.168"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.569 17.9a30 30 0 0 1 10.489 8.495"/>`),
		g.Group(children),
	)
}