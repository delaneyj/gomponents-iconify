package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTxt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileTxt0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" stroke-linejoin="round" d="M10 4h20l10 10v28a2 2 0 0 1-2 2H10a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2Z"/><path stroke="#000" stroke-linecap="round" d="M18 18.008h12m-5.992 0V34"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileTxt0)"/>`),
		g.Group(children),
	)
}