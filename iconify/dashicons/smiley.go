package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smiley(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 5.2a2 2 0 0 1 2 2c0 .37-.11.71-.28 1C8.72 8.2 8 8 7 8s-1.72.2-1.72.2c-.17-.29-.28-.63-.28-1a2 2 0 0 1 2-2zm6 0c1.11 0 2 .89 2 2c0 .37-.11.71-.28 1c0 0-.72-.2-1.72-.2s-1.72.2-1.72.2c-.17-.29-.28-.63-.28-1c0-1.11.89-2 2-2zm-3 13.7a8.69 8.69 0 0 0 8.23-5.88l-1.32-.46C15.9 15.52 13.12 17.5 10 17.5s-5.9-1.98-6.91-4.94l-1.32.46A8.69 8.69 0 0 0 10 18.9z"/>`),
		g.Group(children),
	)
}