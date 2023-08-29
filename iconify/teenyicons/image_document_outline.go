package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageDocumentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 3.5h.5v-.207l-.146-.147l-.354.354Zm-3-3l.354-.354L10.707 0H10.5v.5Zm1 7.995l.354-.353l-.353-.354l-.354.354l.353.353Zm-3 2.998l-.39.313l.349.436l.394-.395l-.353-.354ZM4.5 6.5l.39-.313l-.403-.503L4.1 6.2l.4.3Zm8 7.5h-10v1h10v-1ZM2 13.5v-12H1v12h1Zm11-10v10h1v-10h-1ZM2.5 1h8V0h-8v1Zm7.646-.146l3 3l.708-.708l-3-3l-.708.708ZM2.5 14a.5.5 0 0 1-.5-.5H1A1.5 1.5 0 0 0 2.5 15v-1Zm10 1a1.5 1.5 0 0 0 1.5-1.5h-1a.5.5 0 0 1-.5.5v1ZM2 1.5a.5.5 0 0 1 .5-.5V0A1.5 1.5 0 0 0 1 1.5h1ZM10 6h1V5h-1v1Zm3.854 4.147l-2-2.005l-.708.707l2 2.004l.708-.706Zm-2.707-2.005l-3 2.998l.706.707l3-2.998l-.706-.707ZM8.89 11.18l-4-4.994l-.78.626l4 4.993l.78-.625ZM4.1 6.2l-3 4l.8.6l3-4l-.8-.6Z"/>`),
		g.Group(children),
	)
}