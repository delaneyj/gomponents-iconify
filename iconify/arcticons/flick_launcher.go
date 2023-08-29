package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickLauncher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.735 21.553h8.531v15.4h-8.531zm12.11 11.799h3.613v-8.197h-3.613m-15.69 0h-3.613v8.197h3.613"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.547 18.357L26.562 6.767a4.015 4.015 0 0 0-5.124 0L7.453 18.357A5.397 5.397 0 0 0 5.5 22.513v14.905a4.738 4.738 0 0 0 4.738 4.738h27.524a4.738 4.738 0 0 0 4.738-4.738V22.513a5.4 5.4 0 0 0-1.953-4.156Z"/>`),
		g.Group(children),
	)
}