package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ivonaenglish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" d="M33.614 36.43h-6.08V26.338a3.534 3.534 0 0 0-1.474-2.873l-8.639-6.195a3.535 3.535 0 1 0-4.12 5.744l7.164 5.138v8.278h-6.079a3.535 3.535 0 0 0 0 7.07h19.228a3.535 3.535 0 0 0 0-7.07Z"/><circle cx="27.27" cy="11.917" r="7.417" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.854 11.917h14.833m-7.417 7.416V4.5m-5.962 11.816l11.924-8.799m-11.924 0l11.924 8.799"/>`),
		g.Group(children),
	)
}