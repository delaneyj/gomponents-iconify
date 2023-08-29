package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FreeDownloadManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.675 14.75h18.5c4.084 0 7.4 3.316 7.4 7.4V42.5h-18.5a7.403 7.403 0 0 1-7.4-7.4V14.75h0Zm0 13.875h17.736"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m28.669 33.367l4.742-4.742l-4.742-4.742m-19.82 7.444a7.381 7.381 0 0 1-2.424-5.477V5.5h18.5a7.384 7.384 0 0 1 5.682 2.659"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.643 36.101a7.383 7.383 0 0 1-2.593-5.626v-20.35h18.5a7.38 7.38 0 0 1 5.233 2.167m4.2 4.243l-8.376-8.376M18.46 40.915l-9.6-9.6"/>`),
		g.Group(children),
	)
}