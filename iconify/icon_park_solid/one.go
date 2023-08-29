package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func One(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSOne0"><path fill="#fff" stroke="#fff" stroke-width="4" d="m11 40l-4.06-8.798a5 5 0 0 1 2.1-6.46l11.257-6.29a3 3 0 0 1 3.264.218l1.026.77a1 1 0 0 0 1.595-.697L27.37 7.25a3 3 0 0 1 1.11-2.034l.215-.172a1.865 1.865 0 0 1 2.484.138c.525.524.82 1.236.82 1.978v22.26a1 1 0 0 0 1.624.782L37 27.5c1.653-1.322 3.875-.459 5.255.445c.439.287.504.881.192 1.303L34.5 40S31 44 23 44s-11.333-2.667-12-4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSOne0)"/>`),
		g.Group(children),
	)
}