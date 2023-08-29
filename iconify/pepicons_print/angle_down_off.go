package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDownOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m12 11.547l-5.04-4.2c-1.536-1.28-3.457 1.025-1.92 2.305l5.948 4.957c.272.256.63.42 1.012.418c.382.002.74-.162 1.012-.418l5.948-4.957c1.537-1.28-.384-3.585-1.92-2.304l-5.04 4.2Z" clip-rule="evenodd" opacity=".2"/><path d="m15.68 7.116l-6 5l.64.768l6-5l-.64-.768Z"/><path d="m16.32 7.884l-6 5c-.512.427-1.152-.341-.64-.768l6-5c.512-.427 1.152.341.64.768Z"/><path d="m3.68 7.884l6 5l.64-.768l-6-5l-.64.768Z"/><path d="m4.32 7.116l6 5c.512.427-.128 1.195-.64.768l-6-5c-.512-.427.128-1.195.64-.768ZM1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}