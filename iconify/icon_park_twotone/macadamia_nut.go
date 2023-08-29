package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacadamiaNut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMacadamiaNut0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 43c10.493 0 19-8.059 19-18h-8l-3 3l-3-3H5c0 9.941 8.507 18 19 18Z"/><path d="M40.945 25c.124-.815 0-1.65 0-2.5c0-9.113-7.363-17.5-16.447-17.5c-3.822 0-8.337 1.309-11.13 3.504A4.99 4.99 0 0 1 14.033 11a4.986 4.986 0 0 1-6.004 4.895c-.63 1.75-.974 4.637-.974 6.605c0 .85-.124 1.685 0 2.5M6 31h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMacadamiaNut0)"/>`),
		g.Group(children),
	)
}