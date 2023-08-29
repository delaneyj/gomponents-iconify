package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationOctagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5zm9.854-5.288l-5.566-5.566A.5.5 0 0 0 15.935 2h-7.87a.5.5 0 0 0-.353.146L2.146 7.712A.5.5 0 0 0 2 8.065v7.87a.5.5 0 0 0 .146.353l5.566 5.566a.5.5 0 0 0 .353.146h7.87a.5.5 0 0 0 .353-.146l5.566-5.566a.5.5 0 0 0 .146-.353v-7.87a.5.5 0 0 0-.146-.353zM21 15.728L15.728 21H8.272L3 15.728V8.272L8.272 3h7.456L21 8.272v7.456zm-9-1.353a.625.625 0 1 0 0 1.25a.625.625 0 0 0 0-1.25z"/>`),
		g.Group(children),
	)
}