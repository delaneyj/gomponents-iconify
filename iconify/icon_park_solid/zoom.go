package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSZoom0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M20 43c-4.726-1.767-8.668-7.815-10.64-11.357c-.852-1.53-.403-3.408.964-4.502a3.83 3.83 0 0 1 5.1.283L17 29v-8.5a2.5 2.5 0 0 1 5 0v-4a2.5 2.5 0 0 1 5 0v8a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1 5 0v7.868c0 1.07-.265 2.128-.882 3.003C35.095 39.82 33.255 42.034 31 43c-3.5 1.5-6.63 1.634-11 0Z"/><path d="M13 8h22m-18 4l-4-4l4-4m14 0l4 4l-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSZoom0)"/>`),
		g.Group(children),
	)
}