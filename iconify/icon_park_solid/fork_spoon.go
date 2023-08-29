package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkSpoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSForkSpoon0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M14 4v40M8 5v10c0 5 6 5 6 5s6 0 6-5V5m14 15v24"/><path fill="#fff" d="M40 12c0 4.418-2.686 8-6 8s-6-3.582-6-8s2.686-8 6-8s6 3.582 6 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSForkSpoon0)"/>`),
		g.Group(children),
	)
}