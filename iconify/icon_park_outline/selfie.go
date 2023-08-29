package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selfie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M29.234 6.98c3.714-1.568 11.276-2.47 13.794 5.891c2.518 8.362-8.307 14.612-8.026 8.099c.28-6.514 7.995-1.152 7.83 3.562c-.165 4.714-5.063 5.232-5.063 5.232"/><rect width="26" height="12" x="4.241" y="15.778" rx="2" transform="rotate(-30 4.241 15.778)"/><path stroke-linecap="round" d="M21.5 19.67L29 32.66"/><rect width="6" height="10" x="26.401" y="34.16" rx="3" transform="rotate(-30 26.401 34.16)"/><path stroke-linecap="round" d="m11.401 8.18l5.197-3m5.303 21.186l5.197-3"/></g>`),
		g.Group(children),
	)
}