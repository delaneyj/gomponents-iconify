package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 7.5c-.3 0-.5.2-.5.5v3.5h-5V8c0-.3-.2-.5-.5-.5s-.5.2-.5.5v8c0 .3.2.5.5.5s.5-.2.5-.5v-3.5h5V16c0 .3.2.5.5.5s.5-.2.5-.5V8c0-.3-.2-.5-.5-.5zM12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2zm0 19c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9z"/>`),
		g.Group(children),
	)
}