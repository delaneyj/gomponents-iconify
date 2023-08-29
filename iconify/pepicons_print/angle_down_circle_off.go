package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDownCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="m15 14.547l-5.04-4.2c-1.536-1.28-3.457 1.025-1.92 2.305l5.948 4.957c.272.256.63.42 1.012.418c.382.002.74-.162 1.012-.418l5.948-4.957c1.537-1.28-.384-3.585-1.92-2.304l-5.04 4.2Z" clip-rule="evenodd" opacity=".2"/><path d="m18.68 10.116l-6 5l.64.768l6-5l-.64-.768Z"/><path d="m19.32 10.884l-6 5c-.512.427-1.152-.341-.64-.768l6-5c.512-.427 1.152.341.64.768Z"/><path d="m6.68 10.884l6 5l.64-.768l-6-5l-.64.768Z"/><path d="m7.32 10.116l6 5c.512.427-.128 1.195-.64.768l-6-5c-.512-.427.128-1.195.64-.768ZM4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}