package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmCloudSubnets(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 22a3.607 3.607 0 0 0-2 .6L19.414 18L18 19.414L22.6 24a4.176 4.176 0 0 0-.4 1H9.858a3.553 3.553 0 0 0-.458-1L24 9.4a3.607 3.607 0 0 0 2 .6a4 4 0 1 0-3.857-5H9.9A4.079 4.079 0 0 0 6 2a4 4 0 0 0 0 8a3.607 3.607 0 0 0 2-.6l4.586 4.6L14 12.586L9.4 8a4.175 4.175 0 0 0 .4-1h12.342a3.555 3.555 0 0 0 .458 1L8 22.6a3.607 3.607 0 0 0-2-.6a4 4 0 1 0 3.857 5H22.1a4.012 4.012 0 1 0 3.9-5Zm0-18a2 2 0 1 1-2 2a2.006 2.006 0 0 1 2-2ZM6 8a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Zm0 20a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Zm20 0a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}