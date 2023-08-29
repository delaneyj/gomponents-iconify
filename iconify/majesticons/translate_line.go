package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TranslateLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M9 2a1 1 0 0 1 1 1v1h5a1 1 0 1 1 0 2h-1.546c-.222 2.38-1.212 4.52-2.5 6.341l.827 1.034a1 1 0 1 1-1.562 1.25l-.541-.677c-2.068 2.338-4.539 4.056-6.212 4.937a1 1 0 1 1-.932-1.77c1.49-.784 3.765-2.363 5.654-4.502c.074-.083.147-.168.22-.253l-2.19-2.735a1 1 0 0 1 1.562-1.25l1.865 2.332C10.581 9.27 11.25 7.686 11.442 6H3a1 1 0 0 1 0-2h5V3a1 1 0 0 1 1-1zm7 8a1 1 0 0 1 .894.553l5 10a1 1 0 1 1-1.788.894L18.882 19h-5.764l-1.224 2.447a1 1 0 1 1-1.788-.894l5-10A1 1 0 0 1 16 10zm-1.882 7h3.764L16 13.236L14.118 17z"/></g>`),
		g.Group(children),
	)
}