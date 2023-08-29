package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WiFiRouterMinimalisticBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.11 7.434a4.752 4.752 0 0 1 8.78 0a.75.75 0 0 0 1.386-.574a6.252 6.252 0 0 0-11.553 0a.75.75 0 0 0 1.386.574Z"/><path fill-rule="evenodd" d="M2.586 12.336C2 12.922 2 13.864 2 15.75c0 1.886 0 2.828.586 3.414c.586.586 1.528.586 3.414.586h12c1.886 0 2.828 0 3.414-.586c.586-.586.586-1.528.586-3.414c0-1.886 0-2.828-.586-3.414c-.586-.586-1.528-.586-3.414-.586h-.75v-3a.75.75 0 0 0-1.5 0v3H6c-1.886 0-2.828 0-3.414.586ZM6 16.75a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/><path d="M14.34 8.618a2.251 2.251 0 0 1 4.32 0a.75.75 0 1 0 1.44-.42a3.751 3.751 0 0 0-7.2 0a.75.75 0 0 0 1.44.42Z"/></g>`),
		g.Group(children),
	)
}