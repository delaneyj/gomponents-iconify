package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stepback(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10.5 12.7c-.27 0-.55-.08-.79-.23L4.6 9.27c-.44-.28-.71-.75-.71-1.27s.26-1 .71-1.27l.26.42l-.26-.42l5.11-3.2c.47-.29 1.04-.31 1.52-.04s.77.76.77 1.31v6.39c0 .55-.29 1.04-.77 1.31c-.23.13-.48.19-.73.19Zm0-8.4c-.08 0-.17.02-.26.08l-5.11 3.2c-.21.13-.24.34-.24.42s.02.29.24.42l5.11 3.2c.23.14.43.06.51.01c.08-.04.26-.17.26-.44V4.8a.496.496 0 0 0-.51-.5Zm-9 8.2c-.28 0-.5-.22-.5-.5V4c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}