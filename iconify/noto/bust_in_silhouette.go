package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BustInSilhouette(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#1976D2" d="M115.1 120v4H12.9v-4c0-15.7 19.9-23.3 42-24.9V94c-9.3-2.8-17.6-9.8-21.7-21.3c-4.2-1.5-6.6-15.3-5.5-17.1C26.9 49.8 21.4 4.3 64 4c42.5.2 37.1 45.6 36.2 51.5c1.1 1.8-1.3 15.6-5.5 17.1c-4 11.5-12.3 18.6-21.6 21.4v1c22.2 1.8 42 10.2 42 25z"/>`),
		g.Group(children),
	)
}