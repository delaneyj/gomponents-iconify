package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 2a.75.75 0 0 0-1.5 0v.278c-4.984.38-8.92 4.505-8.999 9.567a.917.917 0 0 0 .766.918c2.727.455 5.479.7 8.233.739V19a.75.75 0 0 1-1.5 0v-.5a.75.75 0 0 0-1.5 0v.5a2.25 2.25 0 0 0 4.5 0v-5.498a54.635 54.635 0 0 0 8.233-.739a.917.917 0 0 0 .766-.918a9.753 9.753 0 0 0-8.999-9.567V2Zm7.476 9.366a8.25 8.25 0 0 0-16.452 0c5.45.854 11.001.854 16.452 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}