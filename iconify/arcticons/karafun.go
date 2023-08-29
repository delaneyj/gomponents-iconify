package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Karafun(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.991 13.308v20.944m11.523-20.607l-8.327 12.448h-3.196m19.85 8.159c-2.145 2.145-3.028 2.44-4.794 2.44c-4.178 0-6.813-10.598-11.86-10.598m-5.028-12.786h3.663m7.608 0h3.663M14.159 34.252h3.663"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}