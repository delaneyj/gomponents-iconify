package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Animeworld(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.933 21.029l-1.831 7.322l-1.83-7.322l-1.831 7.322l-1.831-7.322m-1.215 7.322l-2.38-7.322l-2.471 7.322m.824-2.471h3.203"/><rect width="23.761" height="18.104" x="12.358" y="15.639" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.101"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.294 33.742v4.833m-5.157 0H29.74"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".928" d="M14.357 44.369a19.235 19.235 0 1 1 19.235 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".928" d="M13.742 44.022a21.498 21.498 0 1 1 19.85.347"/>`),
		g.Group(children),
	)
}