package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Obico(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.433 42.804c.3-6.087.23-13.313-3.287-15.766a1.28 1.28 0 0 1-.03-.284a7.41 7.41 0 0 0 3.344-6.169v-.51a7.368 7.368 0 0 0-2.139-5.18a7.212 7.212 0 0 0 1.611-4.27a17.581 17.581 0 0 1-1.98 1.663a15.143 15.143 0 0 1-3.013 1.686c-4.996 2.053-10.55.964-14.952-3.348a7.223 7.223 0 0 0 1.694 4.37a7.363 7.363 0 0 0-2.041 5.079v.51a7.384 7.384 0 0 0 3.424 6.224c0 .12-.01.245-.028.373l-.035.044c-3.93 2.984-3.718 11.657.703 17.337h0C8.902 41.872 2.5 33.683 2.5 24C2.5 12.126 12.126 2.5 24 2.5S45.5 12.126 45.5 24c0 8.088-4.466 15.133-11.067 18.804Z"/>`),
		g.Group(children),
	)
}