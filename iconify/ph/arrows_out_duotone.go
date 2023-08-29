package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 48v48l-48-48ZM48 208h48l-48-48Zm160 0v-48l-48 48ZM48 96l48-48H48Z" opacity=".2"/><path d="M98.34 146.34L72 172.69l-18.34-18.35A8 8 0 0 0 40 160v48a8 8 0 0 0 8 8h48a8 8 0 0 0 5.66-13.66L83.31 184l26.35-26.34a8 8 0 0 0-11.32-11.32ZM56 200v-20.69L76.69 200ZM83.31 72l18.35-18.34A8 8 0 0 0 96 40H48a8 8 0 0 0-8 8v48a8 8 0 0 0 13.66 5.66L72 83.31l26.34 26.35a8 8 0 0 0 11.32-11.32ZM56 76.69V56h20.69ZM208 40h-48a8 8 0 0 0-5.66 13.66L172.69 72l-26.35 26.34a8 8 0 0 0 11.32 11.32L184 83.31l18.34 18.35A8 8 0 0 0 216 96V48a8 8 0 0 0-8-8Zm-8 36.69L179.31 56H200Zm11.06 75.92a8 8 0 0 0-8.72 1.73L184 172.69l-26.34-26.35a8 8 0 0 0-11.32 11.32L172.69 184l-18.35 18.34A8 8 0 0 0 160 216h48a8 8 0 0 0 8-8v-48a8 8 0 0 0-4.94-7.39ZM200 200h-20.69L200 179.31Z"/></g>`),
		g.Group(children),
	)
}