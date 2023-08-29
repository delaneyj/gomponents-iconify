package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemovePlaylist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M465.5 0H186.2c-25.7 0-46.5 20.9-46.5 46.5v279.3c0 25.7 20.9 46.5 46.5 46.5h279.3c25.7 0 46.5-20.9 46.5-46.5V46.5C512 20.9 491.1 0 465.5 0zm-23.3 209.5H209.5V163h232.7v46.5zM93.1 325.8V139.6H46.5C20.9 139.6 0 160.5 0 186.2v279.3C0 491.1 20.9 512 46.5 512h279.3c25.7 0 46.5-20.9 46.5-46.5V419H186.2c-51.4-.1-93.1-41.8-93.1-93.2z"/>`),
		g.Group(children),
	)
}