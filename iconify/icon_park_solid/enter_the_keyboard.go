package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnterTheKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEnterTheKeyboard0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M42 7H6a2 2 0 0 0-2 2v28a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2Z"/><path stroke="#000" stroke-linecap="round" d="M12 19h2m7 0h2m6 0h7m-24 9h24"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEnterTheKeyboard0)"/>`),
		g.Group(children),
	)
}