package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="#F13770" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="#F13770" rx=".6"/><path fill="#FF5C41" d="M2.252 12.688V9.344h27.875v3.344H2.252Z"/><path fill="#FCA34D" d="M2.252 16v-3.344h27.875V16H2.252Z"/><path fill="#43D195" d="M2.252 19.313v-3.344h27.875v3.344H2.252Z"/><path fill="#3F8DF1" d="M2.252 22.656v-3.343h27.875v3.343H2.252Z"/><path fill="#6D549F" d="M2.252 22.656v2.701a.6.6 0 0 0 .6.6h26.675a.6.6 0 0 0 .6-.6v-2.7H2.252Z"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2053id0)" fill-opacity=".25" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2053id1)" fill-opacity=".5" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2053id2)" fill-opacity=".5" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2053id3)" rx=".6"/><defs><linearGradient id="f2053id0" x1="2.252" x2="3.221" y1="17.813" y2="17.813" gradientUnits="userSpaceOnUse"><stop stop-color="#3A3A3A"/><stop offset="1" stop-color="#3A3A3A" stop-opacity="0"/></linearGradient><linearGradient id="f2053id1" x1="30.127" x2="29.159" y1="19.332" y2="19.332" gradientUnits="userSpaceOnUse"><stop stop-color="#FBF2FF"/><stop offset="1" stop-color="#FBF2FF" stop-opacity="0"/></linearGradient><linearGradient id="f2053id2" x1="25.065" x2="25.065" y1="6.043" y2="6.754" gradientUnits="userSpaceOnUse"><stop stop-color="#FBF2FF"/><stop offset="1" stop-color="#FBF2FF" stop-opacity="0"/></linearGradient><linearGradient id="f2053id3" x1="8.752" x2="8.752" y1="26.004" y2="24.938" gradientUnits="userSpaceOnUse"><stop offset=".015" stop-color="#693CA6"/><stop offset="1" stop-color="#693CA6" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}