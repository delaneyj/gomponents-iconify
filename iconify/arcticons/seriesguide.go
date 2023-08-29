package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seriesguide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.22 10a2.33 2.33 0 0 0-1.25.9c-.51.78-.48 0-.46 13.23v12l.23.49A2.44 2.44 0 0 0 6.12 38l.43.18h34.9l.43-.18a2.44 2.44 0 0 0 1.36-1.28l.23-.49v-12c0-13.28 0-12.45-.46-13.23a2.84 2.84 0 0 0-.6-.6c-.77-.48.51-.45-18.46-.45c-16.46 0-17.37 0-17.75.15ZM35 16.87L20.36 31.51s-7.24-7.17-7.36-7.17"/>`),
		g.Group(children),
	)
}