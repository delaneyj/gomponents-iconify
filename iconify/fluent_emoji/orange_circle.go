package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f1532id0)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f1532id4)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f1532id1)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f1532id2)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f1532id3)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f1532id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 9.936 9.935) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop offset=".232" stop-color="#FF8C46"/><stop offset="1" stop-color="#EE534A"/></radialGradient><radialGradient id="f1532id1" cx="0" cy="0" r="1" gradientTransform="rotate(136.38 10.067 10.264) scale(14.6767 15.816)" gradientUnits="userSpaceOnUse"><stop offset=".179" stop-color="#FFA35F"/><stop offset="1" stop-color="#FF8544" stop-opacity="0"/></radialGradient><radialGradient id="f1532id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.249 16)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#E77049" stop-opacity="0"/><stop offset=".951" stop-color="#FFA693"/></radialGradient><radialGradient id="f1532id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 21 -23.3208 0 15.757 9)" gradientUnits="userSpaceOnUse"><stop offset=".863" stop-color="#EE4C57" stop-opacity="0"/><stop offset="1" stop-color="#EC5B84"/></radialGradient><linearGradient id="f1532id4" x1="15.757" x2="15.757" y1="2" y2="8.5" gradientUnits="userSpaceOnUse"><stop stop-color="#FF7E51"/><stop offset="1" stop-color="#EF4B5E" stop-opacity="0"/><stop offset="1" stop-color="#FF7E51" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}