package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f133id0)" d="M29.547 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f133id4)" d="M29.547 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f133id1)" d="M29.547 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f133id2)" d="M29.547 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f133id3)" d="M29.547 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f133id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 9.831 9.886) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop stop-color="#645B70"/><stop offset=".662" stop-color="#33293D"/></radialGradient><radialGradient id="f133id1" cx="0" cy="0" r="1" gradientTransform="matrix(-9.47663 8.74999 -8.07914 -8.75007 24.227 10.75)" gradientUnits="userSpaceOnUse"><stop offset=".113" stop-color="#6A6571"/><stop offset="1" stop-color="#49474C" stop-opacity="0"/></radialGradient><radialGradient id="f133id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.04 16)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#434344" stop-opacity="0"/><stop offset=".951" stop-color="#8C8B8D"/></radialGradient><radialGradient id="f133id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 21.8837 -24.3022 0 15.547 9)" gradientUnits="userSpaceOnUse"><stop offset=".853" stop-color="#352343" stop-opacity="0"/><stop offset="1" stop-color="#50405F"/></radialGradient><linearGradient id="f133id4" x1="15.547" x2="15.547" y1="2" y2="8.5" gradientUnits="userSpaceOnUse"><stop stop-color="#554A63"/><stop offset="1" stop-color="#554A63" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}