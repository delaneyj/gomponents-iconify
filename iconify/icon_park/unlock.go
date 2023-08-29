package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="34" height="22" x="7" y="22.048" fill="#2F88FF" stroke="#000" rx="2"/><path stroke="#000" stroke-linecap="round" d="M14 22V14.0047C13.9948 8.87022 17.9227 4.56718 23.0859 4.05117C28.249 3.53516 32.9673 6.97408 34 12.0059"/><path stroke="#fff" stroke-linecap="round" d="M24 30V36"/></g>`),
		g.Group(children),
	)
}