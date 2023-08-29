package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diesix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 959H128q-53 0-90.5-37.5T0 831V127q0-53 37.5-90T128 0h704q53 0 90.5 37t37.5 90v704q0 53-37.5 90.5T832 959zM224 831q40 0 68-28t28-68t-28-68t-68-28t-68 28t-28 68t28 68t68 28zm0-704q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5zm0 256q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5zm512-256q-40 0-68 28.5t-28 68t28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5zm-96 352.5q0 39.5 28 67.5t68 28t68-28t28-67.5t-28-68t-68-28.5t-68 28.5t-28 68zM736 639q-40 0-68 28t-28 68t28 68t68 28t68-28t28-68t-28-68t-68-28z"/>`),
		g.Group(children),
	)
}