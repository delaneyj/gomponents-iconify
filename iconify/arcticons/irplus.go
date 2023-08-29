package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Irplus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="31.2" height="39" x="8.4" y="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.95"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.4 8.5h3.61m-3.61 4.82h3.61m-3.61 4.82h3.61M18.6 8.5h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m3.6-9.64h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m3.59-9.64h3.61m-3.61 4.82h3.61m-3.61 4.82h3.61m-25.2 6.9h3.61m-3.61 4.82h3.61m-3.61 4.82h3.61M11.4 39.5h3.61m3.59-14.46h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m3.6-14.46h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m-3.6 4.82h3.6m3.59-14.46h3.61m-3.61 4.82h3.61m-3.61 4.82h3.61m-3.61 4.82h3.61"/>`),
		g.Group(children),
	)
}