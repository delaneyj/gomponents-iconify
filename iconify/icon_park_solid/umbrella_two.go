package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUmbrellaTwo0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M27 24v13.125C27 39 26.638 44 23 44c-3.429 0-4-4.375-4-5.625"/><path fill="#fff" stroke="#fff" d="M24 4c14.5 0 19.375 13.333 20 20H4c.625-6.667 5.5-20 20-20Z"/><path stroke="#000" d="m19 14l4 4l6-7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUmbrellaTwo0)"/>`),
		g.Group(children),
	)
}