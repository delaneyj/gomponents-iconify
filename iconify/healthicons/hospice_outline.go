package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospiceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M41.708 23.794L24.058 6.086L6.294 23.792a1 1 0 1 0 1.412 1.416L10 22.922V41a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V22.907l2.292 2.3a1 1 0 1 0 1.416-1.413ZM24.053 8.914L36 20.9V40H12V20.928L24.053 8.914Z" clip-rule="evenodd"/><path d="M16 25.517C16 22.655 17.794 20 20.25 20c1.704 0 2.97 1.191 3.75 2.88c.78-1.689 2.045-2.88 3.75-2.88c2.456 0 4.25 2.655 4.25 5.517C32 31.645 24 36 24 36s-5.389-2.756-7.32-7.278a8.863 8.863 0 0 1-.36-1a7.851 7.851 0 0 1-.32-2.205Z"/></g>`),
		g.Group(children),
	)
}