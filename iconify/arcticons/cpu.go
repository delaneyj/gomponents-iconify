package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.014 5.5h17.971v3.171H15.014zm0 33.829h17.971V42.5H15.014zM42.5 15.015v17.971h-3.171V15.015zm-37-.001h3.171v17.971H5.5z"/><rect width="20" height="20" x="14" y="13.347" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.671 8.671h30.657v30.657H8.671z"/>`),
		g.Group(children),
	)
}