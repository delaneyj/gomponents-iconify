package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.155 28.832L19.168 7.845a8.007 8.007 0 0 0-11.323 0h0a8.007 8.007 0 0 0 0 11.323l20.987 20.987a8.007 8.007 0 0 0 11.323 0h0a8.007 8.007 0 0 0 0-11.323Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.323 24l4.832-4.832a8.007 8.007 0 0 0 0-11.323h0a8.007 8.007 0 0 0-11.323 0l-4.847 4.847M12.677 24l-4.832 4.832a8.007 8.007 0 0 0 0 11.323h0a8.007 8.007 0 0 0 11.323 0l4.817-4.817"/><circle cx="34.345" cy="34.494" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}