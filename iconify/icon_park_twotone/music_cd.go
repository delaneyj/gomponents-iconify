package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicCd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMusicCd0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M26 14v14"/><path fill="#555" stroke-linejoin="round" d="M14 28.666C14 26.64 15.934 25 18.32 25H26v4.334C26 31.36 24.066 33 21.68 33h-3.36C15.934 33 14 31.359 14 29.334v-.668Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m32 15l-6-1"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMusicCd0)"/>`),
		g.Group(children),
	)
}