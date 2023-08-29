package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerPhoneOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 22h-4q-.825 0-1.413-.588T8 20v-8q0-.825.588-1.413T10 10h4q.825 0 1.413.588T16 12v8q0 .825-.588 1.413T14 22Zm0-2v-8h-4v8h4Zm0 0h-4h4ZM7.725 7.775q-.3-.3-.3-.713t.35-.662q.925-.675 2-1.038T12 5q1.15 0 2.225.363t2 1.037q.35.25.363.65t-.313.725q-.275.275-.7.3t-.8-.225q-.6-.425-1.312-.637T12 7q-.75 0-1.463.213t-1.312.637q-.375.25-.8.225t-.7-.3ZM4.95 4.9q-.3-.3-.288-.7t.313-.65q1.5-1.25 3.288-1.9T12 1q1.95 0 3.738.65t3.287 1.9q.3.25.313.65t-.288.7q-.275.275-.688.3t-.762-.25Q16.4 4 14.962 3.5T12 3q-1.525 0-2.963.5T6.4 4.95q-.35.275-.763.25t-.687-.3Z"/>`),
		g.Group(children),
	)
}