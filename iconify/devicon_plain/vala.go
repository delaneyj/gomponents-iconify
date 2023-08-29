package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vala(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M62.959 32.051c-7.551 0-14.777 2.427-21.676 7.283c-3.02 2.191-5.491 4.708-7.416 7.551c-1.895 2.813-2.844 5.745-2.844 8.795c0 1.658.209 3.049.623 4.174c1.214 3.109 4.248 4.664 9.104 4.664c0-.178-.073-.384-.22-.621c-.682-1.244-1.022-3.095-1.022-5.553c0-5.063 1.183-9.298 3.552-12.703c2.399-3.405 5.73-5.907 9.994-7.506l1.465 62.137h13.102l25.273-67.777h-6.351L67.4 88.192l-.801-55.963a36.416 36.416 0 0 0-3.64-.178z"/>`),
		g.Group(children),
	)
}