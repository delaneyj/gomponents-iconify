package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthereumOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5.5l.384-.32a.5.5 0 0 0-.768 0L7.5.5Zm-5 6l-.384-.32a.5.5 0 0 0-.04.585L2.5 6.5Zm5 8l-.424.265a.5.5 0 0 0 .848 0L7.5 14.5Zm5-8l.424.265a.5.5 0 0 0-.04-.585l-.384.32Zm-5-2l.186-.464L7.5 3.96l-.186.075l.186.464ZM7.116.18l-5 6l.768.64l5-6l-.768-.64Zm-5.04 6.585l5 8l.848-.53l-5-8l-.848.53Zm5.848 8l5-8l-.848-.53l-5 8l.848.53Zm4.96-8.585l-5-6l-.768.64l5 6l.768-.64Zm-10.198.784l5-2l-.372-.928l-5 2l.372.928Zm4.628-2l5 2l.372-.928l-5-2l-.372.928Z"/>`),
		g.Group(children),
	)
}