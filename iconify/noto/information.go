package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Information(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#427687" d="M116.46 3.96h-104c-4.42 0-8 3.58-8 8v104c0 4.42 3.58 8 8 8h104c4.42 0 8-3.58 8-8v-104c0-4.42-3.58-8-8-8z"/><path fill="#8CAFBF" d="M110.16 3.96h-98.2a7.555 7.555 0 0 0-7.5 7.5v97.9c-.01 4.14 3.34 7.49 7.48 7.5h98.12c4.14.01 7.49-3.34 7.5-7.48V11.46c.09-4.05-3.13-7.41-7.18-7.5h-.22z"/><path fill="#FFF" d="M67.76 108.96h-11.9a2.433 2.433 0 0 1-2.4-2.4v-56.8a2.433 2.433 0 0 1 2.4-2.4h11.8c1.32.02 2.38 1.08 2.4 2.4v56.8c.03 1.3-1 2.37-2.3 2.4z"/><circle cx="61.76" cy="29.66" r="10.7" fill="#FFF"/><path fill="#B4E1ED" d="M40.16 12.86c0-2.3-1.6-3-10.8-2.7c-7.7.3-11.5 1.2-13.8 4s-2.9 8.5-3 15.3c0 4.8 0 9.3 2.5 9.3c3.4 0 3.4-7.9 6.2-12.3c5.4-8.7 18.9-10.6 18.9-13.6z" opacity=".5"/>`),
		g.Group(children),
	)
}