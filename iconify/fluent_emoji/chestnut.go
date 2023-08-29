package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chestnut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#D88578" d="M29.68 22.172L16 17L2.32 22.172c1.04 3.97 4.52 7.28 9.39 7.28h8.57c4.88 0 8.36-3.31 9.4-7.28Z"/><path fill="url(#f244id0)" d="M29.68 22.172L16 17L2.32 22.172c1.04 3.97 4.52 7.28 9.39 7.28h8.57c4.88 0 8.36-3.31 9.4-7.28Z"/><path fill="url(#f244id1)" d="M29.68 22.172L16 17L2.32 22.172c1.04 3.97 4.52 7.28 9.39 7.28h8.57c4.88 0 8.36-3.31 9.4-7.28Z"/><path fill="url(#f244id4)" d="M29.68 22.172c.83-3.16.11-6.75-2.69-9.42l-9.59-9.19c-.78-.75-2.02-.75-2.8 0l-9.59 9.19c-2.8 2.68-3.51 6.26-2.69 9.42h27.36Z"/><path fill="url(#f244id2)" d="M29.68 22.172c.83-3.16.11-6.75-2.69-9.42l-9.59-9.19c-.78-.75-2.02-.75-2.8 0l-9.59 9.19c-2.8 2.68-3.51 6.26-2.69 9.42h27.36Z"/><path fill="url(#f244id3)" d="M29.68 22.172c.83-3.16.11-6.75-2.69-9.42l-9.59-9.19c-.78-.75-2.02-.75-2.8 0l-9.59 9.19c-2.8 2.68-3.51 6.26-2.69 9.42h27.36Z"/><defs><radialGradient id="f244id0" cx="0" cy="0" r="1" gradientTransform="matrix(0 10 -21.9715 0 16 30.75)" gradientUnits="userSpaceOnUse"><stop stop-color="#CC6B7E"/><stop offset=".632" stop-color="#CD6C75" stop-opacity="0"/></radialGradient><radialGradient id="f244id1" cx="0" cy="0" r="1" gradientTransform="matrix(0 7.2025 -17.75 0 25.75 22.25)" gradientUnits="userSpaceOnUse"><stop offset=".087" stop-color="#E7A277"/><stop offset=".694" stop-color="#ECA27A" stop-opacity="0"/></radialGradient><radialGradient id="f244id2" cx="0" cy="0" r="1" gradientTransform="matrix(.125 7.25 -14.51844 .25033 23.5 18.25)" gradientUnits="userSpaceOnUse"><stop stop-color="#9A7772"/><stop offset=".509" stop-color="#916E69" stop-opacity=".4"/><stop offset="1" stop-color="#744450" stop-opacity="0"/></radialGradient><radialGradient id="f244id3" cx="0" cy="0" r="1" gradientTransform="matrix(.62492 10.50002 -24.15316 1.4375 12.626 24)" gradientUnits="userSpaceOnUse"><stop stop-color="#693848"/><stop offset="1" stop-color="#7B4552" stop-opacity="0"/></radialGradient><linearGradient id="f244id4" x1="24.125" x2="2" y1="9.875" y2="18.25" gradientUnits="userSpaceOnUse"><stop stop-color="#835C55"/><stop offset=".549" stop-color="#643E38"/><stop offset=".871" stop-color="#5A403C"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}