package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberSmartRecordRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19q-2.925 0-4.963-2.038T2 12q0-2.925 2.038-4.963T9 5q2.925 0 4.963 2.038T16 12q0 2.925-2.038 4.963T9 19Zm8.5-.475q-.575.2-1.038-.038T16 17.6q0-.275.188-.525t.462-.35q1.5-.5 2.425-1.8T20 12q0-1.625-.925-2.925t-2.425-1.8q-.275-.1-.463-.35T16 6.4q0-.625.45-.875t1.025-.05Q19.5 6.2 20.75 7.988T22 12q0 2.225-1.25 4t-3.25 2.525Z"/>`),
		g.Group(children),
	)
}