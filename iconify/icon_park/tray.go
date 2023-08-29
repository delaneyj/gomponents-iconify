package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="28" height="30" x="24.762" y="3.243" fill="#2F88FF" stroke="#000" rx="2" transform="rotate(45 24.762 3.243)"/><path stroke="#000" d="M38.1966 16.6775L42.4392 12.4348L35.3682 5.36377L31.1255 9.60641"/><path stroke="#fff" d="M18 21H30"/><path stroke="#fff" d="M18 27H22"/><path stroke="#fff" d="M28 27H30"/></g>`),
		g.Group(children),
	)
}