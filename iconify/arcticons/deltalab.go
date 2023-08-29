package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deltalab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24c-.086 11.205-5.827 13.018-.259 20.899c-8.855-3.993-10.937.645-21.241.6A21.5 21.5 0 1 1 45.5 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.387 9.628c1.09 1.66 5.314 5.433 6.324 3.375c1.404-2.86-7.53-3.724-10.68-3.607c-3.266.121-6.087 2.637-1.755 6.424c3.978 3.477 9.863 3.181 12.42 7.28c9.769 15.66-16.269 21.164-17.517 8.523c-1.087-11.005 13.194-11.854 16.926-9.33"/>`),
		g.Group(children),
	)
}