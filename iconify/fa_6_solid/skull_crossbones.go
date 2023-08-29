package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkullCrossbones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M368 128c0 44.4-25.4 83.5-64 106.4V256c0 17.7-14.3 32-32 32h-96c-17.7 0-32-14.3-32-32v-21.6c-38.6-23-64-62.1-64-106.4C80 57.3 144.5 0 224 0s144 57.3 144 128zm-200 48a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm144-32a32 32 0 1 0-64 0a32 32 0 1 0 64 0zM3.4 273.7c7.9-15.8 27.1-22.2 42.9-14.3L224 348.2l177.7-88.8c15.8-7.9 35-1.5 42.9 14.3s1.5 35-14.3 42.9L295.6 384l134.8 67.4c15.8 7.9 22.2 27.1 14.3 42.9s-27.1 22.2-42.9 14.3L224 419.8L46.3 508.6c-15.8 7.9-35 1.5-42.9-14.3s-1.5-35 14.3-42.9L152.4 384L17.7 316.6c-15.8-7.9-22.2-27.1-14.3-42.9z"/>`),
		g.Group(children),
	)
}