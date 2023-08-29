package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Two(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<linearGradient id="svgIDa" x1="64.002" x2="64.002" y1="102" y2="26.24" gradientTransform="matrix(1 0 0 -1 0 128)" gradientUnits="userSpaceOnUse"><stop offset="0" stop-color="#757575"/><stop offset=".515" stop-color="#504F4F"/></linearGradient><path fill="url(#svgIDa)" d="M87.86 102H41.51c-1.12 0-2.03-.91-2.03-2.03v-5.78c0-.51.19-1 .53-1.37l24.44-26.66c3.6-4.02 6.17-7.37 7.7-10.07c1.53-2.69 2.29-5.4 2.29-8.11c0-3.57-1-6.45-3.01-8.65s-4.71-3.3-8.11-3.3c-4.05 0-7.19 1.24-9.42 3.71c-1.89 2.1-2.98 4.85-3.27 8.25c-.09 1.06-.95 1.89-2.02 1.89h-8.46c-1.18 0-2.12-1-2.03-2.17c.28-3.71 1.29-7.09 3.03-10.16c2.08-3.66 5.05-6.5 8.91-8.52S58.38 26 63.42 26c7.28 0 13.02 1.83 17.22 5.48c4.2 3.66 6.31 8.71 6.31 15.16c0 3.74-1.06 7.66-3.17 11.77c-2.11 4.1-5.57 8.76-10.38 13.98L58.27 88.6c-1.21 1.3-.29 3.41 1.48 3.41h28.11c1.12 0 2.03.91 2.03 2.03v5.93c0 1.12-.91 2.03-2.03 2.03z"/>`),
		g.Group(children),
	)
}