package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crying(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13z"/><path fill="currentColor" d="M12.5 6h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1zm-7 0h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1zm4 7.375a.502.502 0 0 1-.354-.146C9.074 13.157 8.686 13 8 13s-1.075.157-1.146.229a.5.5 0 0 1-.707-.707c.471-.471 1.453-.521 1.854-.521s1.383.051 1.854.521a.5.5 0 0 1-.354.853zM11.5 9a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5zm0 3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5zm-7-3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5zm0 3a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 1 0v1a.5.5 0 0 1-.5.5z"/>`),
		g.Group(children),
	)
}