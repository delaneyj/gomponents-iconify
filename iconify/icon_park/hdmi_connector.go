package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdmiConnector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 43C29 40.2386 26.7614 38 24 38C21.2386 38 19 40.2386 19 43"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 42.1091C18.0797 43.3221 20.9607 44 24 44C27.0393 44 29.9203 43.3221 32.5 42.1091"/><circle cx="15" cy="15" r="3" fill="#fff"/><circle cx="11" cy="23" r="3" fill="#fff"/><circle cx="24" cy="11" r="3" fill="#fff"/><circle cx="33" cy="15" r="3" fill="#fff"/><circle cx="37" cy="23" r="3" fill="#fff"/></g>`),
		g.Group(children),
	)
}