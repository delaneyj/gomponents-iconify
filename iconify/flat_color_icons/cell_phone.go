package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellPhone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#546E7A" d="M12 40V10h20c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4z"/><path fill="#4FC3F7" d="M32 13H16c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h16c.6 0 1-.4 1-1v-8c0-.6-.4-1-1-1z"/><path fill="#B3E5FC" d="M19 30h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm-12 5h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm-12 5h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1zm6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1z"/><path fill="#37474F" d="M16 10h-4V4c0-1.1.9-2 2-2s2 .9 2 2v6z"/>`),
		g.Group(children),
	)
}