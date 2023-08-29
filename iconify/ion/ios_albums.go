package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosAlbums(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M460.9 161H51.1C31.8 161 16 176.8 16 196.1V428c0 19.3 15.8 35.1 35.1 35.1H461c19.3 0 35.1-15.8 35.1-35.1V196.1c-.1-19.3-15.9-35.1-35.2-35.1z" fill="currentColor"/><path d="M434 133H78c-7.7 0-14-6.3-14-14s6.3-14 14-14h356c7.7 0 14 6.3 14 14s-6.3 14-14 14z" fill="currentColor"/><path d="M403.2 77H108.8c-7 0-12.8-5.8-12.8-12.8v-2.4c0-7 5.8-12.8 12.8-12.8h294.4c7 0 12.8 5.8 12.8 12.8v2.4c0 7-5.8 12.8-12.8 12.8z" fill="currentColor"/>`),
		g.Group(children),
	)
}