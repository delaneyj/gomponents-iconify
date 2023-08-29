package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pasta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m896.59 320l-257-52q-5 49-41 82.5t-86 33.5h-64V64h64q50 0 86 33.5t41 81.5l257-51q128 0 128 95q0 53-36 75t-92 22zm-544 128q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5t22.5 9.5t9.5 22.5v384q0 13-9.5 22.5t-22.5 9.5zm-128 576q-13 0-22.5-9.5t-9.5-22.5V32q0-13 9.5-22.5t22.5-9.5t22.5 9.5t9.5 22.5v960q0 13-9.5 22.5t-22.5 9.5zM.59 352q0-13 9.5-22.5t22.5-9.5h96v64h-96q-13 0-22.5-9.5T.59 352zm0-128q0-13 9.5-22.5t22.5-9.5h96v64h-96q-13 0-22.5-9.5T.59 224zm0-128q0-13 9.5-22.5t22.5-9.5h96v64h-96q-13 0-22.5-9.5T.59 96z"/>`),
		g.Group(children),
	)
}