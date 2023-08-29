package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loudspeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 1.382V15.5l-7.2-3h-.91L7.712 14H9.5v2H7.477l-.588 5H2V5h6.764L16 1.382ZM5.875 12.5H4V19h1.11l.765-6.5ZM4 10.5h5.2l4.8 2V4.618L9.236 7H4v3.5Zm16.333-5.914l.707.707c.937.937 1.328 2.342 1.328 3.624c0 1.281-.391 2.687-1.328 3.624l-.707.707l-1.414-1.414l.707-.708c.444-.443.742-1.264.742-2.21c0-.945-.298-1.765-.742-2.209L18.919 6l1.414-1.414ZM18.25 6.255l.707.707a2.768 2.768 0 0 1 0 3.914l-.707.707l-1.415-1.414l.707-.707c.3-.3.3-.786 0-1.086l-.707-.707l1.415-1.414Z"/>`),
		g.Group(children),
	)
}