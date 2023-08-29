package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnePieceSwimsuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="m23.114 22.5l3.5 16.5l-6.5 14.5c7.447 3.468 11.748 8.946 13.175 15.5h4.65c.561-6.332 5.664-11.282 13.175-15.5l-6.5-14.5l4.05-16.5c-1.55 6-22.55 7-25.55 0Z"/><path fill="#8967aa" d="M44.886 4.5s6.612 18.22 2 21c-3.996 2.408-16.822 2.849-22 0c-5.877-3.234 2-21 2-21c-.289.29-2.04 4.044 2 12c6.15 6.258 8.129 6.87 13 0c5.205-9.105 3.191-11.858 3-12Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m23.114 22.5l3.5 16.5l-6.5 14.5c7.447 3.468 11.748 8.946 13.175 15.5h4.65c.561-6.332 5.664-11.282 13.175-15.5l-6.5-14.5l4.05-16.5"/><path d="M45.122 4.5s6.612 18.22 2 21c-3.996 2.408-16.822 2.849-22 0c-5.878-3.234 2-21 2-21s-3.431 4.222 2 12c4.194 6.007 9.225 6.447 13 0c4.976-8.5 3-12 3-12Z"/></g>`),
		g.Group(children),
	)
}