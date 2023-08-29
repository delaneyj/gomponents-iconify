package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Npxs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#F5D100"/><g fill="#FFF"><path d="M15.972 4C9.372 4 4 9.372 4 15.972s5.372 11.982 11.982 11.982c6.609 0 11.981-5.372 11.981-11.982C27.963 9.363 22.582 4 15.973 4zm0 23.03c-6.092 0-11.058-4.956-11.058-11.058C4.924 9.88 9.88 4.923 15.972 4.923c6.093 0 11.059 4.957 11.059 11.059c0 6.092-4.957 11.049-11.059 11.049z"/><path d="m9.774 10.656l.88-.881l4.688 4.687l-.881.88zm6.818 6.821l.881-.881l4.687 4.687l-.881.88zm-1.495-1.502l.88-.88l.888.886l-.88.882zm1.508-1.475l4.687-4.686l.88.881l-4.686 4.687zm-6.797 6.788l4.686-4.686l.881.88l-4.686 4.687z"/></g></g>`),
		g.Group(children),
	)
}