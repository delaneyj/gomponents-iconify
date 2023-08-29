package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibilityNewRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm8 2.25q-1.2.275-2.475.463T15 9v12.025q0 .425-.288.7T14 22q-.425 0-.713-.288T13 21v-5h-2v5.025q0 .425-.288.7T10 22q-.425 0-.713-.288T9 21V9q-1.25-.1-2.525-.288T4 8.25q-.425-.1-.65-.463T3.25 7q.1-.425.463-.638T4.5 6.25q1.8.4 3.713.575T12 7q1.875 0 3.788-.175T19.5 6.25q.425-.1.775.113T20.75 7q.1.425-.113.788T20 8.25Z"/>`),
		g.Group(children),
	)
}