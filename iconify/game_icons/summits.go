package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Summits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96.07 20c-26.51 0-48 21.49-48 48c0 26.5 21.49 48 48 48c26.53 0 48.03-21.5 48.03-48c0-26.51-21.5-48-48.03-48zM326.2 81.5L217.1 237.9l-6.5 69.2l-86.1-108.7l-68.48 111.4l-10.63 56.1l-29.26 67.7l1.13.5L64.07 492l96.03-21.3l25.3-25.4l-69.5-93.5l45.4-24.8l-34.5-96.5l136.3 180.2l9-14.7l-19.1-86l66.8-28.6l.3-156.4l51.7 178.7l-44.6 62.8l-9 39.4l-50.8 54.8L403 426.6l-37.1-21.2l34.5-31.8l-27.8-23.3l52.8-72.3l6.1 90.6l52 49.9l12.4-13l-47-45.1l-7.5-112.8l-57.3 33.8z"/>`),
		g.Group(children),
	)
}