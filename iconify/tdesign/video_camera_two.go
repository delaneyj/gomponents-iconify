package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 3a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7Zm4.243 7a5.5 5.5 0 1 0-9.45-5.278A4 4 0 0 0 1.536 10H1v12h17v-2.523l5 2V10.523l-5 2V10h-1.257ZM18 14.677l3-1.2v5.046l-3-1.2v-2.646ZM16 12v8H3v-8h13ZM5 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}