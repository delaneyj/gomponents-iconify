package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbUpRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21q-.825 0-1.413-.588T8 19V8.825q0-.4.163-.762t.437-.638l5.425-5.4q.375-.35.888-.425t.987.175q.475.25.688.7t.087.925L15.55 8H21q.8 0 1.4.6T23 10v2q0 .175-.038.375t-.112.375l-3 7.05q-.225.5-.75.85T18 21h-8Zm-6 0q-.825 0-1.413-.588T2 19v-9q0-.825.588-1.413T4 8q.825 0 1.413.588T6 10v9q0 .825-.588 1.413T4 21Z"/>`),
		g.Group(children),
	)
}