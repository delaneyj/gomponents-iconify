package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cambio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.395 15.914c1.33-.208 2.684-.248 4.083 0l2.042-2.042c.205-1.042-.567-1.76-2.042-2.042c-1.113-.212-2.3-.523-3.145-.228c-3.356 1.354-8.973 8.606-8.83 14.472c.07 2.875 2.086 8.12 7.258 9.838c3.113 1.034 8.08.44 11.863.418c4.724-.026 10.319-2.457 15.128-4.302c.284-.894.774-.934 4.748-4.107c-1.126-.36-4.532.354-6.443 1.05c-5.097 1.859-10.731 2.956-17.574 2.931c-4.329-.015-10.384-1.268-11.385-5.731c-1.017-4.536-.055-7.7 4.297-10.258Z"/>`),
		g.Group(children),
	)
}