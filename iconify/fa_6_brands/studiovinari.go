package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Studiovinari(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="m480.3 187.7l4.2 28v28l-25.1 44.1l-39.8 78.4l-56.1 67.5l-79.1 37.8l-17.7 24.5l-7.7 12l-9.6 4s17.3-63.6 19.4-63.6c2.1 0 20.3.7 20.3.7l66.7-38.6l-92.5 26.1l-55.9 36.8l-22.8 28l-6.6 1.4l20.8-73.6l6.9-5.5l20.7 12.9l88.3-45.2l56.8-51.5l14.8-68.4l-125.4 23.3l15.2-18.2l-173.4-53.3l81.9-10.5l-166-122.9L133.5 108L32.2 0l252.9 126.6l-31.5-38L378 163L234.7 64l18.7 38.4l-49.6-18.1L158.3 0l194.6 122L310 66.2l108 96.4l12-8.9l-21-16.4l4.2-37.8L451 89.1l29.2 24.7l11.5 4.2l-7 6.2l8.5 12l-13.1 7.4l-10.3 20.2l10.5 23.9z"/>`),
		g.Group(children),
	)
}