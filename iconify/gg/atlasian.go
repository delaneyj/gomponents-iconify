package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Atlasian(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.507 11.556c-.317-.452-.725-.397-.911.122L5 18.908h5.178c.52-2.058.07-4.865-1.097-6.533l-.574-.819Z" opacity=".8"/><path d="M12.874 7.126c-1.267 1.81-1.675 4.958-.911 7.03l1.75 4.751h5.251l-4.597-12.48c-.19-.519-.602-.572-.919-.12l-.574.82Z"/></g>`),
		g.Group(children),
	)
}