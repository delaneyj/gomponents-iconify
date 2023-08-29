package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BustInSilhouette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.34 3.07h-2.69A5 5 0 0 0 8.7 8.73l.078 1.204a2 2 0 0 0 .257 3.975l.115 1.761a4.75 4.75 0 0 0 3.597 3.988A12.92 12.92 0 0 0 6.19 22.93v-.01a12.91 12.91 0 0 0-3.5 5.53v.11a2.995 2.995 0 0 0-.09.32l3.587.04l.002.01h17.62l.003-.01l3.588-.04l-.022-.059a.766.766 0 0 1-.068-.261v-.11a13 13 0 0 0-3.5-5.53v.01a12.922 12.922 0 0 0-6.552-3.27a4.749 4.749 0 0 0 3.602-3.99l.109-1.764a2 2 0 0 0 .245-3.961l.076-1.215a4.999 4.999 0 0 0-4.95-5.66Z"/>`),
		g.Group(children),
	)
}