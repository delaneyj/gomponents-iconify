package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 16c-4.455 0-6.685-5.386-3.535-8.535C9.615 4.315 15 6.545 15 11a5 5 0 0 1-5 5zM6.42 2.56l-.67.64c-.37.357-.865.808-1.38.81H2C.914 4 0 4.712 0 5.76v10.48C0 17.27 1 18 2 18h16c1 0 2-.716 2-1.76V5.76C20 4.723 19 4 18 4h-2.37c-.515-.002-1.01-.453-1.38-.81l-.67-.64A2 2 0 0 0 12.2 2H7.8a2 2 0 0 0-1.38.56z"/><circle cx="10" cy="11" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}