package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphoneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3.75C8.436 3.75 5.75 6.205 5.75 9v1.261c.082-.007.166-.011.25-.011h3a.75.75 0 0 1 .75.75v7a.75.75 0 0 1-.75.75H6A2.75 2.75 0 0 1 3.25 16v-3c0-.854.39-1.617 1-2.121V9c0-3.832 3.582-6.75 7.75-6.75S19.75 5.168 19.75 9v1.879c.61.504 1 1.267 1 2.121v3A2.75 2.75 0 0 1 18 18.75h-3a.75.75 0 0 1-.75-.75v-7a.75.75 0 0 1 .75-.75h3c.084 0 .168.004.25.011V9c0-2.795-2.686-5.25-6.25-5.25Zm-6 8c-.69 0-1.25.56-1.25 1.25v3c0 .69.56 1.25 1.25 1.25h2.25v-5.5H6ZM19.25 13c0-.69-.56-1.25-1.25-1.25h-2.25v5.5H18c.69 0 1.25-.56 1.25-1.25v-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}