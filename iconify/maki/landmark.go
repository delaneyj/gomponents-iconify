package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.5 12H12v-.5c0-.3-.2-.5-.5-.5H11V6h1l1-2c-1 .1-2 .1-3 0c-.8-.6-1.4-1.2-2-2v-.5c0-.3-.2-.5-.5-.5s-.5.2-.5.5V2c-.6.8-1.2 1.4-2 2c-1 .1-2 .1-3 0l1 2h1v5h-.5c-.3 0-.5.2-.5.5v.5h-.5c-.3 0-.5.2-.5.5v.5h11v-.5c0-.3-.2-.5-.5-.5zM7 11H5V6h2v5zm3 0H8V6h2v5z"/>`),
		g.Group(children),
	)
}