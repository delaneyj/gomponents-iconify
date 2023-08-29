package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationsActiveRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.425 0-.713-.288T4 18q0-.425.288-.713T5 17h1v-7q0-2.075 1.25-3.688T10.5 4.2v-.7q0-.625.438-1.063T12 2q.625 0 1.063.438T13.5 3.5v.7q2 .5 3.25 2.113T18 10v7h1q.425 0 .713.288T20 18q0 .425-.288.713T19 19H5Zm7 3q-.825 0-1.413-.588T10 20h4q0 .825-.588 1.413T12 22ZM3 10q-.425 0-.713-.325t-.237-.75q.2-1.875 1.05-3.488t2.175-2.812q.325-.275.738-.25t.662.375q.25.35.2.75t-.375.7q-.975.925-1.6 2.15T4.075 9q-.05.425-.35.713T3 10Zm18 0q-.425 0-.725-.288T19.925 9q-.2-1.425-.825-2.65T17.5 4.2q-.325-.3-.375-.7t.2-.75q.25-.35.663-.375t.737.25q1.325 1.2 2.175 2.812t1.05 3.488q.05.425-.237.75T21 10Z"/>`),
		g.Group(children),
	)
}