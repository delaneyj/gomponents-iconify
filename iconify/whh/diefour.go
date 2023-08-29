package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diefour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 959H128q-53 0-90.5-37.5T0 831V127q0-52 37.5-89.5T128 0h704q53 0 90.5 37t37.5 90v704q0 53-37.5 90.5T832 959zM224 127q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5zm0 512q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28zm512-512q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5zm0 512q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}