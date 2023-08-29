package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwitchOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5h.5Zm14 0h.5a.5.5 0 0 0-.5-.5v.5Zm0 8l.354.354A.5.5 0 0 0 15 8.5h-.5Zm-3 3v.5a.5.5 0 0 0 .354-.146L11.5 11.5Zm-5 0V11a.5.5 0 0 0-.325.12l.325.38Zm-3.5 3h-.5a.5.5 0 0 0 .825.38L3 14.5Zm0-3h.5A.5.5 0 0 0 3 11v.5Zm-2.5 0H0a.5.5 0 0 0 .5.5v-.5ZM.5 1h14V0H.5v1ZM14 .5v8h1v-8h-1Zm.146 7.646l-3 3l.708.708l3-3l-.708-.708ZM11.5 11h-5v1h5v-1Zm-5.325.12l-3.5 3l.65.76l3.5-3l-.65-.76ZM3.5 14.5v-3h-1v3h1ZM3 11H.5v1H3v-1Zm-2 .5V.5H0v11h1ZM10 3v5h1V3h-1ZM7 3v5h1V3H7Z"/>`),
		g.Group(children),
	)
}