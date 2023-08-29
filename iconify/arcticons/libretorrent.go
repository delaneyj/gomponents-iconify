package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libretorrent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.289 7.314A19.009 19.009 0 0 0 6.283 35.637m34.789-14.105A19.028 19.028 0 0 0 26.833 7.314M12.272 41.808a19.008 19.008 0 0 0 28.806-11.703"/><circle cx="22.561" cy="6.832" r="4.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="41.561" cy="25.832" r="4.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.278" cy="38.721" r="4.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.243 34.886c7.228-6.757 13.616-9.63 24.018-9.332"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.547 11.132c-.09 7.61-1.005 15.314-10.148 24.634"/>`),
		g.Group(children),
	)
}