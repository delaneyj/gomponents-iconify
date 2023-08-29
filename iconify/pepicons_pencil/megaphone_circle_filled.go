package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilMegaphoneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.4 7.882v4.436l7.6 2.073V5.809L9.4 7.882Zm-1 4.436a1 1 0 0 0 .737.965l7.6 2.072A1 1 0 0 0 18 14.391V5.809a1 1 0 0 0-1.263-.965l-7.6 2.073a1 1 0 0 0-.737.965v4.436Z"/><path d="M6.456 12.3H8.5V7.9H6.453a3.422 3.422 0 0 0 .003 4.4Zm2.044 1a1 1 0 0 0 1-1V7.9a1 1 0 0 0-1-1H6.253a.55.55 0 0 0-.4.172c-1.602 1.691-1.595 4.353-.002 6.052a.555.555 0 0 0 .405.176H8.5Z"/><path d="m10.269 13.87l-2.51-.28l-.978 3.91h2.636l.852-3.63Zm-2.4-1.273a1 1 0 0 0-1.081.75l-.977 3.91a1 1 0 0 0 .97 1.243h2.636a1 1 0 0 0 .974-.772l.852-3.63a1 1 0 0 0-.864-1.223l-2.51-.278Zm13.747-6.374a.5.5 0 0 1-.139.693l-1.5 1a.5.5 0 1 1-.554-.832l1.5-1a.5.5 0 0 1 .693.139ZM19.2 10.1a.5.5 0 0 1 .5-.5h1.5a.5.5 0 0 1 0 1h-1.5a.5.5 0 0 1-.5-.5Zm.117 2.23a.5.5 0 0 1 .705-.06l1.38 1.159a.5.5 0 0 1-.643.765l-1.38-1.16a.5.5 0 0 1-.062-.704Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilMegaphoneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}