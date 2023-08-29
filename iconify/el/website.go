package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Website(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M698.598 922.244h-199.6V722.646h199.6v199.598zm252.506 0H751.503V722.646h199.601v199.598zm-505.012 0H246.493V722.646h199.599v199.598zm505.012-631.262v367.936h-704.61V290.982h704.61zm101-105.812H147.896v829.66h904.209V185.17h-.001zM0 1161.521V38.478h1200v1123.045l-1200-.002z"/>`),
		g.Group(children),
	)
}