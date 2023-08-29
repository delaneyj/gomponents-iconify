package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoaderOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M8 1V.5H7V1h1ZM7 4.5V5h1v-.5H7Zm1 6V10H7v.5h1ZM7 14v.5h1V14H7ZM4.5 7.995H5v-1h-.5v1Zm-3.5-1H.5v1H1v-1ZM14 8h.5V7H14v1Zm-3.5-1.005H10v1h.5v-1ZM7 1v3.5h1V1H7Zm0 9.5V14h1v-3.5H7ZM4.5 6.995H1v1h3.5v-1ZM14 7l-3.5-.005v1L14 8V7ZM2.147 2.854l3 3l.708-.708l-3-3l-.708.708Zm10-.708l-3 3l.708.708l3-3l-.708-.708ZM2.854 12.854l3-3l-.708-.708l-3 3l.708.708Zm6.292-3l3 3l.708-.708l-3-3l-.708.708Z"/>`),
		g.Group(children),
	)
}