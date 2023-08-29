package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnseenGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm7.442 27.512l18.023-18.024"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.703 19.817a4.476 4.476 0 0 0-5.853 5.853m2.805 2.585a4.485 4.485 0 0 0 5.627-5.63"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.49 17.03A17.656 17.656 0 0 0 9.5 23.94a17.828 17.828 0 0 0 6.494 5.586m3.95 1.44A17.694 17.694 0 0 0 38.5 23.941a17.859 17.859 0 0 0-6.17-5.407"/>`),
		g.Group(children),
	)
}