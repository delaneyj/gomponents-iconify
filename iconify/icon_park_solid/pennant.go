package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pennant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPennant0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 44h4M12 6V4v2Zm0 16v22v-22Zm0 22H8h4Zm-4 0h8"/><path fill="#fff" d="M12 6v16l28-8l-28-8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPennant0)"/>`),
		g.Group(children),
	)
}