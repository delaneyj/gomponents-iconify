package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiPeopleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm-2 16q-.425 0-.713-.288T9 21V8.775Q7 8.25 5.712 6.688T4.1 3.124q-.075-.45.238-.788T5.15 2q.35 0 .613.25t.312.6q.275 1.775 1.55 2.963T10.75 7h2.5q.75 0 1.4.275t1.175.8L19.65 11.9q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L15 10.05V21q0 .425-.288.713T14 22q-.425 0-.713-.288T13 21v-5h-2v5q0 .425-.288.713T10 22Z"/>`),
		g.Group(children),
	)
}