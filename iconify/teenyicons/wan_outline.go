package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WanOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m3.5 14.5l-.489.105a.5.5 0 0 0 .967.042L3.5 14.5Zm4-13l.478-.147a.5.5 0 0 0-.956 0L7.5 1.5Zm4 13l-.478.147a.5.5 0 0 0 .967-.042L11.5 14.5ZM.011.605l3 14l.978-.21l-3-14l-.978.21Zm3.967 14.042l4-13l-.956-.294l-4 13l.956.294Zm3.044-13l4 13l.956-.294l-4-13l-.956.294Zm4.967 12.958l3-14l-.978-.21l-3 14l.978.21ZM0 6h15V5H0v1Zm0 4h15V9H0v1Z"/>`),
		g.Group(children),
	)
}