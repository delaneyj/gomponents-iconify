package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSyncDismissTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.414 3.635a.5.5 0 0 0 0-.707L9.293.807a.5.5 0 0 0-.707.707l.997.997a7.5 7.5 0 0 0-4.075 13.495a.5.5 0 0 0 .6-.8A6.5 6.5 0 0 1 10.066 3.5c.43.004.901.069 1.322.16l.025-.025ZM8.586 16.363l.025-.025c.42.092.892.156 1.323.16a6.5 6.5 0 0 0 3.959-11.706a.5.5 0 1 1 .6-.8a7.5 7.5 0 0 1-4.075 13.495l.996.997a.5.5 0 1 1-.707.707l-2.121-2.12a.5.5 0 0 1 0-.708ZM15 9.999a5 5 0 1 1-10 0a5 5 0 0 1 10 0Zm-3.147-1.146a.5.5 0 0 0-.707-.707L10 9.292L8.853 8.146a.5.5 0 0 0-.707.707l1.147 1.146l-1.147 1.147a.5.5 0 1 0 .707.707L10 10.706l1.146 1.147a.5.5 0 0 0 .707-.707l-1.146-1.147l1.146-1.146Z"/>`),
		g.Group(children),
	)
}