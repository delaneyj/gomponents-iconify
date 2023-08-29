package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowFrameBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M3.464 3.464c-1.08 1.08-1.363 2.647-1.438 5.286h19.948c-.075-2.64-.358-4.205-1.439-5.286C19.072 2 16.714 2 12 2S4.929 2 3.464 3.464ZM2 12c0-.623 0-1.206.003-1.75H8.25v11.706c-2.34-.097-3.775-.41-4.786-1.42C2 19.07 2 16.713 2 12Zm4-5a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/><path d="M22 12c0 4.714 0 7.071-1.465 8.535C19.072 22 16.714 22 12 22c-.819 0-1.566 0-2.25-.008V10.25h12.247C22 10.794 22 11.377 22 12Z"/></g>`),
		g.Group(children),
	)
}