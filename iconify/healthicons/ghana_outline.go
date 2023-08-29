package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GhanaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M24 7a1 1 0 0 0-1 1v2.98l-.857.124a13.003 13.003 0 0 0-8.952 5.644a13 13 0 0 0 8.952 20.089l.857.124V40a1 1 0 1 0 2 0v-3.04l.857-.123a13 13 0 0 0 7.335-3.674a1 1 0 0 0-1.414-1.414a11 11 0 0 1-5.578 3l-1.2.244V12.948l1.2.245a11 11 0 0 1 5.578 3a1 1 0 0 0 1.414-1.415a13 13 0 0 0-7.335-3.674L25 10.98V8a1 1 0 0 0-1-1Zm-3 1a3 3 0 1 1 6 0v1.273a15 15 0 0 1 7.607 4.09a3 3 0 1 1-4.243 4.243A9 9 0 0 0 27 15.486v16.97a9 9 0 0 0 3.364-2.122a3 3 0 0 1 4.243 4.243A15 15 0 0 1 27 38.667V40a3 3 0 1 1-6 0v-1.333a15.007 15.007 0 0 1-9.472-6.363A15 15 0 0 1 21 9.274V8Zm2 4.948v22.045l-1.2-.245a11 11 0 0 1 0-21.555l1.2-.245Zm-2 2.537a9 9 0 0 0 0 16.97v-16.97Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}