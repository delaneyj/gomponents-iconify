package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransgenderFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2399id0)" rx=".6"/><path fill="#FB8190" d="M2.252 22.125V9.875h27.875v12.25H2.252Z"/><path fill="#FCECFF" d="M2.252 18.063v-4.125h27.875v4.124H2.252Z"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2399id1)" fill-opacity=".25" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2399id2)" fill-opacity=".5" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2399id3)" fill-opacity=".5" rx=".6"/><rect width="27.875" height="19.914" x="2.252" y="6.043" fill="url(#f2399id4)" rx=".6"/><defs><linearGradient id="f2399id0" x1="9.313" x2="30.127" y1="15.25" y2="15.25" gradientUnits="userSpaceOnUse"><stop stop-color="#39A7F8"/><stop offset="1" stop-color="#54C9FA"/></linearGradient><linearGradient id="f2399id1" x1="2.252" x2="3.221" y1="17.813" y2="17.813" gradientUnits="userSpaceOnUse"><stop stop-color="#3A3A3A"/><stop offset="1" stop-color="#3A3A3A" stop-opacity="0"/></linearGradient><linearGradient id="f2399id2" x1="30.127" x2="29.159" y1="19.332" y2="19.332" gradientUnits="userSpaceOnUse"><stop stop-color="#FBF2FF"/><stop offset="1" stop-color="#FBF2FF" stop-opacity="0"/></linearGradient><linearGradient id="f2399id3" x1="25.065" x2="25.065" y1="6.043" y2="6.754" gradientUnits="userSpaceOnUse"><stop stop-color="#FBF2FF"/><stop offset="1" stop-color="#FBF2FF" stop-opacity="0"/></linearGradient><linearGradient id="f2399id4" x1="8.752" x2="8.752" y1="26.281" y2="25.406" gradientUnits="userSpaceOnUse"><stop offset=".015" stop-color="#693CA6"/><stop offset="1" stop-color="#693CA6" stop-opacity="0"/></linearGradient></defs></g>`),
		g.Group(children),
	)
}