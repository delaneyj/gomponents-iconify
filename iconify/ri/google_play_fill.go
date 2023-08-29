package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.61 1.814L13.793 12L3.611 22.186a.996.996 0 0 1-.61-.92V2.735a1 1 0 0 1 .609-.921ZM14.5 12.707l2.302 2.302l-10.937 6.333l8.635-8.635Zm3.199-3.198l2.807 1.626a1 1 0 0 1 0 1.73l-2.808 1.626L15.207 12l2.492-2.491ZM5.865 2.658L16.803 8.99L14.5 11.293L5.865 2.658Z"/>`),
		g.Group(children),
	)
}