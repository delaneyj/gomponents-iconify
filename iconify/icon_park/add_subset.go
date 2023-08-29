package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddSubset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 28V35H18"/><path fill="#2F88FF" d="M18 28H42V42H18V35V28Z"/><line x1="6" x2="6" y1="13.5" y2="12.5"/><line x1="6" x2="6" y1="20" y2="19"/><line x1="6" x2="6" y1="7" y2="6"/><line x1="32" x2="32" y1="13.5" y2="12.5"/><line x1="32" x2="32" y1="20" y2="19"/><line x1="32" x2="32" y1="7" y2="6"/><line x1="32" x2="31" y1="20" y2="20"/><line x1="7" x2="6" y1="20" y2="20"/><line x1="7" x2="6" y1="6" y2="6"/><line x1="13" x2="12" y1="6" y2="6"/><line x1="19.5" x2="18.5" y1="6" y2="6"/><line x1="19.5" x2="18.5" y1="20" y2="20"/><line x1="26" x2="25" y1="6" y2="6"/><line x1="13" x2="12" y1="20" y2="20"/><line x1="26" x2="25" y1="20" y2="20"/><line x1="32" x2="31" y1="6" y2="6"/></g>`),
		g.Group(children),
	)
}