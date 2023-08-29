package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandleC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="24" r="20" fill="#2F88FF" stroke="#000"/><path stroke="#fff" d="M32 17.6183C31.1017 15.7887 28.4068 12.5867 23.0171 13.0442C17.6273 13.5017 13.5842 18.5332 14.0342 24.937C14.4842 31.3407 19.4239 35 23.9154 35C29.3051 35 32 30.6089 32 30.6089"/></g>`),
		g.Group(children),
	)
}