package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5.421L14 4.213v6.574L7.5 14.58L1 10.787V4.213L7.5.42ZM7.5 4a3.5 3.5 0 1 0 0 7h.182c.976 0 1.913-.388 2.604-1.078l.068-.068l-.708-.708l-.068.068A2.682 2.682 0 0 1 7.682 10H7.5a2.5 2.5 0 0 1 0-5h.182c.711 0 1.393.283 1.896.786l.068.068l.708-.708l-.068-.068A3.682 3.682 0 0 0 7.682 4H7.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}