package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlueCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f145id0)" d="M29.757 15.625c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f145id4)" d="M29.757 15.625c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f145id1)" d="M29.757 15.625c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f145id2)" d="M29.757 15.625c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f145id3)" d="M29.757 15.625c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f145id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 10.023 9.747) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop offset=".116" stop-color="#4DA0DA"/><stop offset=".853" stop-color="#3557C3"/></radialGradient><radialGradient id="f145id1" cx="0" cy="0" r="1" gradientTransform="rotate(136.38 10.142 10.077) scale(14.6767 15.816)" gradientUnits="userSpaceOnUse"><stop offset=".179" stop-color="#62A5E9"/><stop offset="1" stop-color="#4579D7" stop-opacity="0"/></radialGradient><radialGradient id="f145id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.249 15.625)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#416AA9" stop-opacity="0"/><stop offset=".951" stop-color="#859BC6"/></radialGradient><radialGradient id="f145id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 21 -23.3208 0 15.757 8.625)" gradientUnits="userSpaceOnUse"><stop offset=".863" stop-color="#3E4DCB" stop-opacity="0"/><stop offset="1" stop-color="#695FD4"/></radialGradient><linearGradient id="f145id4" x1="15.757" x2="15.757" y1="1.625" y2="8.125" gradientUnits="userSpaceOnUse"><stop stop-color="#5B9CCE"/><stop offset="1" stop-color="#5B9CCE" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}