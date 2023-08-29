package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="url(#f2085id0)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2085id4)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2085id1)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2085id2)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><path fill="url(#f2085id3)" d="M29.757 16c0 7.732-6.268 14-14 14s-14-6.268-14-14s6.268-14 14-14s14 6.268 14 14Z"/><defs><radialGradient id="f2085id0" cx="0" cy="0" r="1" gradientTransform="rotate(130.168 9.936 9.935) scale(27.8086)" gradientUnits="userSpaceOnUse"><stop offset=".232" stop-color="#F24756"/><stop offset="1" stop-color="#B22945"/></radialGradient><radialGradient id="f2085id1" cx="0" cy="0" r="1" gradientTransform="rotate(136.38 10.067 10.264) scale(14.6767 15.816)" gradientUnits="userSpaceOnUse"><stop offset=".179" stop-color="#FF6180"/><stop offset="1" stop-color="#E5364A" stop-opacity="0"/></radialGradient><radialGradient id="f2085id2" cx="0" cy="0" r="1" gradientTransform="matrix(-19.25 0 0 -20 20.249 16)" gradientUnits="userSpaceOnUse"><stop offset=".62" stop-color="#B73E4B" stop-opacity="0"/><stop offset=".951" stop-color="#D48387"/></radialGradient><radialGradient id="f2085id3" cx="0" cy="0" r="1" gradientTransform="matrix(0 21 -23.3208 0 15.757 9)" gradientUnits="userSpaceOnUse"><stop offset=".863" stop-color="#B83C5A" stop-opacity="0"/><stop offset="1" stop-color="#B83C5A"/><stop offset="1" stop-color="#AC4064"/></radialGradient><linearGradient id="f2085id4" x1="15.757" x2="15.757" y1="2" y2="8.5" gradientUnits="userSpaceOnUse"><stop stop-color="#DD4577"/><stop offset="1" stop-color="#EF4B5E" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}