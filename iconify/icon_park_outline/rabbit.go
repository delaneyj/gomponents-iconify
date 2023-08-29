package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rabbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><ellipse cx="24" cy="32" stroke="currentColor" stroke-width="4" rx="17" ry="12"/><circle cx="18" cy="29.412" r="2" fill="currentColor"/><circle cx="24" cy="35.412" r="2" fill="currentColor"/><circle cx="30" cy="29.412" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12.667 22c-1.315-3.167-3.607-10.887-2.254-16.087c.376-1 1.803-2.7 4.509-1.5c.375.166 1.24.8 1.69 2C17.74 8.413 16.05 21 16.05 21m19.337 1c1.316-3.167 3.545-10.887 2.192-16.087c-.376-1-1.804-2.7-4.51-1.5c-.376.166-1.24.8-1.691 2c-1.128 2 .626 13.587.626 13.587"/></g>`),
		g.Group(children),
	)
}