package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ionic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<g fill="#4e8ef7"><circle cx="64" cy="64" r="24.08"/><path d="M113.14 23.14a8.27 8.27 0 0 0-13.7-6.25a59 59 0 1 0 11.67 11.67a8.24 8.24 0 0 0 2.03-5.42zM64 121A57 57 0 1 1 98.1 18.36a8.27 8.27 0 0 0 11.53 11.53A57 57 0 0 1 64 121z"/></g>`),
		g.Group(children),
	)
}