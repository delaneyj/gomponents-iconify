package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trakt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.904 4.433L8.141 29.121L32.904 4.433Zm6.331 34.737L18.699 18.638L39.235 39.17ZM32.757 8.637L22.441 18.979L32.757 8.637Zm.448 3.609l-8.858 8.915l8.858-8.915ZM8.64 32.764l10.034-10.065l18.178 18.178"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.804 37.657l10.932-10.932L34.313 42.55"/>`),
		g.Group(children),
	)
}