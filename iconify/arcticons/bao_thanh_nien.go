package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaoThanhNien(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.3 14.37h4.816m-2.408 7.268V14.37m4.239 0v7.268m4.815-7.268v7.268m-4.815-3.648h4.815m15.123-3.62v7.268M39.7 14.37v7.268m-4.815-3.648H39.7m-18.062 1.217h4.724m0-4.837v7.268m-4.724-.022v-2.41c0-2.724 1.995-4.837 4.724-4.837m6.646 7.269V14.37m-4.724 0v7.268m4.724-2.432c0-2.724-1.995-4.837-4.724-4.837M13.025 33.63v-7.268m-4.725 0v7.268m4.725-2.431c0-2.724-1.995-4.837-4.725-4.837m7.36 0v7.268m3.795-8.794l.654-.363l.654.363m-2.469 5.16h2.37m1.263 3.634h-3.635v-7.268h3.634m7.216 7.268v-7.268m-4.725 0v7.268m4.725-2.431c0-2.724-1.995-4.837-4.725-4.837"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}