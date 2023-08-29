package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TShirtCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.448 13.396L6 13.12V20a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-6.882l.553.276a1 1 0 0 0 1.423-.677l1-4.5a1 1 0 0 0-.269-.924l-2.414-2.414A3 3 0 0 0 18.172 4H16a1 1 0 0 0-.986.83C14.748 6.382 14.129 7 13 7c-1.13 0-1.747-.618-2.013-2.17A1 1 0 0 0 10 4H7.829a3 3 0 0 0-2.122.88L3.295 7.292a1 1 0 0 0-.27.923l.999 4.503a1 1 0 0 0 1.424.678Zm1.673-7.103A1 1 0 0 1 7.83 6h1.377c.57 1.928 1.871 3 3.793 3c1.922 0 3.223-1.072 3.795-3h1.378a1 1 0 0 1 .707.293l2.026 2.026l-.603 2.714l-.855-.427A1 1 0 0 0 18 11.5V20a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-8.5a1 1 0 0 0-1.448-.894l-.853.428l-.602-2.715l2.024-2.026Z" clip-rule="evenodd"/><path d="M14.5 13.5a1 1 0 1 1 0-2h2a1 1 0 1 1 0 2h-2Z"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}