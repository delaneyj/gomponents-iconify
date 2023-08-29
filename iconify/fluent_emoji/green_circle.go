package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f616id0)" d="M29.757 15.75c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f616id4)" d="M29.757 15.75c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f616id1)" d="M29.757 15.75c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f616id2)" d="M29.757 15.75c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f616id3)" d="M29.757 15.75c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f616id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 9.994 9.81) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop offset=".19" stop-color="#5AE68D"/><stop offset=".835" stop-color="#43A684"/></radialGradient><radialGradient id="f616id1" cx="0" cy="0" r="1" gradientTransform="rotate(136.38 10.117 10.14) scale(14.6767 15.816)" gradientUnits="userSpaceOnUse"><stop offset=".179" stop-color="#6FFCA5"/><stop offset="1" stop-color="#65E6A7" stop-opacity="0"/></radialGradient><radialGradient id="f616id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.249 15.75)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#64CB85" stop-opacity="0"/><stop offset=".951" stop-color="#A4E4B7"/></radialGradient><radialGradient id="f616id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 22.1875 -22.9876 0 15.757 8.75)" gradientUnits="userSpaceOnUse"><stop offset=".732" stop-color="#4A9795" stop-opacity="0"/><stop offset="1" stop-color="#718CAD"/></radialGradient><linearGradient id="f616id4" x1="15.757" x2="15.757" y1="1.75" y2="8.25" gradientUnits="userSpaceOnUse"><stop stop-color="#5ED284"/><stop offset="1" stop-color="#5ED284" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}