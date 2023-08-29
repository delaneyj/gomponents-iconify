package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whirlwind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M24 29C26.7614 29 29 26.7614 29 24C29 21.2386 26.7614 19 24 19C21.2386 19 19 21.2386 19 24C19 26.7614 21.2386 29 24 29Z"/><path d="M23.5 44C16.5964 44 11 38.4036 11 31.5C11 24.5964 16.5964 19 23.5 19"/><path d="M44 23.5C44 30.4036 38.4036 36 31.5 36C24.5964 36 19 30.4036 19 23.5"/><path d="M23.5 29C30.4036 29 36 23.4036 36 16.5C36 9.59644 30.4036 4 23.5 4"/><path d="M29 23.5C29 16.5964 23.4036 11 16.5 11C9.59644 11 4 16.5964 4 23.5"/></g>`),
		g.Group(children),
	)
}