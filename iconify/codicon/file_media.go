package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileMedia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 2h6v3.5l.5.5H12v1h1V4.8l-.15-.36l-3.28-3.3L9.22 1H1.5l-.5.5v13l.5.5H5v-1H2V2zm7 0l3 3H9V2zm5.5 6h-8l-.5.5v6l.5.5h8l.5-.5v-6l-.5-.5zM14 9v4l-1.63-1.6h-.71l-1.16 1.17l-2.13-2.13h-.71L7 11.1V9h7zm-2.8 4.27l.81-.81L13.55 14h-1.62l-.73-.73zM7 14v-1.49l1-1L10.52 14H7zm5.5-3.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}