package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Talenthouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.373 7.42V0H1.627v7.42h6.66V24h7.428V7.42h6.66zM12.31 0h-.623zm-.004 3.41V.618h8.865L17.652 3.41Zm-5.948 0L2.83.618h8.857V3.41H6.358zm-.608.308l-3.503 2.76V.949ZM2.837 6.802l3.52-2.781h4.894L8.46 6.8H2.837Zm6.068.438l2.78-2.782v14.781l-1.602 2.046l-1.183 1.51Zm.326 16.142l.555-.706l2.216-2.825l2.77 3.535zm3.078-18.924l2.786 2.782v15.556l-2.786-3.556zM15.55 6.8l-2.8-2.78h4.904l3.519 2.78H15.55Zm6.206-.322L18.25 3.71L21.744.963l.02-.015Z"/>`),
		g.Group(children),
	)
}