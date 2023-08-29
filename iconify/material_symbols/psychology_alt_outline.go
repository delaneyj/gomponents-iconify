package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsychologyAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22v-4.3q-1.425-1.3-2.212-3.038T3 11q0-3.75 2.625-6.375T12 2q3.125 0 5.538 1.838t3.137 4.787l1.3 5.125q.125.475-.175.863T21 15h-2v3q0 .825-.588 1.413T17 20h-2v2h-2v-4h4v-5h2.7l-.95-3.875q-.575-2.275-2.45-3.7T12 4Q9.1 4 7.05 6.025T5 10.95q0 1.5.613 2.85t1.737 2.4l.65.6V22H6Zm6.35-9ZM12 16q.425 0 .713-.288T13 15q0-.425-.288-.713T12 14q-.425 0-.713.288T11 15q0 .425.288.713T12 16Zm-.75-3.2h1.525q0-.625.163-1.012t.662-.938q.45-.5.875-1.012T14.9 8.5q0-1.05-.812-1.775T12.075 6q-1 0-1.812.575t-1.138 1.5l1.375.575q.175-.55.613-.888t.962-.337q.55 0 .913.3t.362.775q0 .525-.312.938t-.738.837q-.5.525-.775 1.05T11.25 12.8Z"/>`),
		g.Group(children),
	)
}