package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaArchive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m12 2l4 4v12H4V2h8zm0 4h3l-3-3v3zM8 3.5v2l1.8-1zM11 5L9.2 6L11 7V5zM8 6.5v2l1.8-1zM11 8L9.2 9l1.8 1V8zM8 9.5v2l1.8-1zm3 1.5l-1.8 1l1.8 1v-2zm-1.5 6c.83 0 1.62-.72 1.5-1.63c-.05-.38-.49-1.61-.49-1.61l-1.99-1.1s-.45 1.95-.52 2.71c-.07.77.67 1.63 1.5 1.63zm0-2.39c.42 0 .76.34.76.76c0 .43-.34.77-.76.77s-.76-.34-.76-.77c0-.42.34-.76.76-.76z"/>`),
		g.Group(children),
	)
}