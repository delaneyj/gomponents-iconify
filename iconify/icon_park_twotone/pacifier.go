package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pacifier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPacifier0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m30.656 21.585l4.103-4.103a6 6 0 1 0-4.243-4.243l-4.102 4.104"/><path fill="#555" d="M41.263 32.192L15.807 6.736l-3.889 3.89l7.778 11.313l-2.828 4.243l4.596 4.596l4.243-2.828l11.667 8.131l3.889-3.889Z"/><path d="M19.32 21.885a9.998 9.998 0 0 0-9.877 2.529c-3.905 3.905-3.905 10.237 0 14.142c3.905 3.905 10.237 3.905 14.142 0a9.998 9.998 0 0 0 2.53-9.877"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPacifier0)"/>`),
		g.Group(children),
	)
}