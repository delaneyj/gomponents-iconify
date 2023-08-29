package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundcloudLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10a1 1 0 0 1 1 1v7a1 1 0 1 1-2 0v-7a1 1 0 0 1 1-1Zm3 1a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0v-6a1 1 0 0 1 1-1Zm3-4a1 1 0 0 1 1 1v10a1 1 0 1 1-2 0V8a1 1 0 0 1 1-1Zm5-1a6 6 0 0 1 5.996 5.775l.003.26a3.5 3.5 0 0 1-.307 6.96L20.5 19h-3.501a1 1 0 0 1-.117-1.993L17 17h3.447l.138-.002a1.5 1.5 0 0 0 .267-2.957l-.135-.026l-1.77-.252l.053-1.787l-.004-.176A4 4 0 0 0 15.2 8.005L15 8c-.268 0-.531.026-.788.077L14 8.126V18a1 1 0 0 1-.883.993L13 19a1 1 0 0 1-1-1V6.803A5.973 5.973 0 0 1 15 6ZM1 12a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}