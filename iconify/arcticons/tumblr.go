package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tumblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.526 43.5c-5.865 0-10.236-3.017-10.236-10.237V21.702h-5.33V15.44c5.865-1.522 8.318-6.57 8.6-10.941h6.092v9.926h7.106v7.276h-7.106v10.067c0 3.017 1.523 4.06 3.947 4.06h3.441v7.67h-6.514Z"/>`),
		g.Group(children),
	)
}