package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrowaveOven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="30" x="4" y="6" stroke="#000" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38.0002 15H34.0002"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 21H34"/><rect width="17" height="12" x="10" y="15" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><circle cx="36" cy="27" r="2" fill="#000"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36V42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 36V42"/></g>`),
		g.Group(children),
	)
}