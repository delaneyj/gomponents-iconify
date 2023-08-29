package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UiZoomOutNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" fill-rule="evenodd" clip-path="url(#healthiconsUiZoomOutNegative0)" clip-rule="evenodd"><path d="M40 21c0 7.18-5.82 13-13 13s-13-5.82-13-13S19.82 8 27 8s13 5.82 13 13Zm-7 1H21v-2h12v2Z"/><path d="M0 0h48v48H0V0Zm12.927 32.245l1.705-.12l1.158-1.158A14.944 14.944 0 0 1 12 21c0-8.284 6.716-15 15-15c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15a14.942 14.942 0 0 1-9.784-3.63l-1.346 1.346l-.158 1.4L8.828 42L6 39.172l6.927-6.927Z"/></g><defs><clipPath id="healthiconsUiZoomOutNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}