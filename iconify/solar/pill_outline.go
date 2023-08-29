package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.132 3.46A7.546 7.546 0 1 0 3.46 14.132l6.408 6.408A7.546 7.546 0 0 0 20.54 9.868L14.132 3.46ZM4.521 4.52a6.046 6.046 0 0 1 8.55 0l2.898 2.898l-.011.035c-.09.282-.243.698-.49 1.21c-.494 1.024-1.365 2.438-2.866 3.94c-1.502 1.5-2.915 2.372-3.94 2.865c-.512.247-.927.4-1.209.49l-.035.01l-2.897-2.896a6.046 6.046 0 0 1 0-8.551ZM8.59 17.14l2.338 2.34a6.046 6.046 0 0 0 8.551-8.552L17.14 8.59c-.088.218-.194.461-.32.724c-.56 1.162-1.526 2.718-3.158 4.349c-1.63 1.631-3.186 2.597-4.348 3.157c-.263.126-.506.232-.724.32Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}