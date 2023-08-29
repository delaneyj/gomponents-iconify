package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YellowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f2964id0)" d="M29.757 15.875c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2964id4)" d="M29.757 15.875c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2964id1)" d="M29.757 15.875c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2964id2)" d="M29.757 15.875c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2964id3)" d="M29.757 15.875c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f2964id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 9.965 9.872) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop offset=".232" stop-color="#F3BB4B"/><stop offset=".959" stop-color="#C79738"/></radialGradient><radialGradient id="f2964id1" cx="0" cy="0" r="1" gradientTransform="rotate(136.38 10.092 10.202) scale(14.6767 15.816)" gradientUnits="userSpaceOnUse"><stop offset=".179" stop-color="#FFE764"/><stop offset="1" stop-color="#F8CA4D" stop-opacity="0"/></radialGradient><radialGradient id="f2964id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.249 15.875)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#C69B40" stop-opacity="0"/><stop offset=".951" stop-color="#E8C290"/></radialGradient><radialGradient id="f2964id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 21 -23.3208 0 15.757 8.875)" gradientUnits="userSpaceOnUse"><stop offset=".792" stop-color="#E5A152" stop-opacity="0"/><stop offset="1" stop-color="#D17887"/></radialGradient><linearGradient id="f2964id4" x1="15.757" x2="15.757" y1="1.875" y2="8.375" gradientUnits="userSpaceOnUse"><stop stop-color="#E1B45D"/><stop offset="1" stop-color="#E1B45D" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}