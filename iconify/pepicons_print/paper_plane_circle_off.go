package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M4.874 12.454a.5.5 0 0 1-.048-.952l15.715-5.855a.5.5 0 0 1 .654.61l-4.744 16.085a.5.5 0 0 1-.95.026L12.696 14.5l-7.823-2.046Z" opacity=".2"/><path fill-rule="evenodd" d="m3.874 10.454l7.823 2.046l2.803 7.868a.5.5 0 0 0 .95-.026l4.745-16.085a.5.5 0 0 0-.654-.61L3.826 9.502a.5.5 0 0 0 .048.952Zm1.783-.567l13.296-4.954l-4.027 13.652l-2.376-6.67a.5.5 0 0 0-.344-.315L5.657 9.887Z" clip-rule="evenodd"/><path d="m19 4.293l.707.707L12 12.707L11.293 12L19 4.293Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}