package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsOutSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 48v48l-48-48ZM48 208h48l-48-48Z" opacity=".2"/><path d="M208 40h-48a8 8 0 0 0-5.66 13.66L172.69 72l-34.35 34.34a8 8 0 0 0 11.32 11.32L184 83.31l18.34 18.35A8 8 0 0 0 216 96V48a8 8 0 0 0-8-8Zm-8 36.69L179.31 56H200Zm-93.66 61.65L72 172.69l-18.34-18.35A8 8 0 0 0 40 160v48a8 8 0 0 0 8 8h48a8 8 0 0 0 5.66-13.66L83.31 184l34.35-34.34a8 8 0 0 0-11.32-11.32ZM56 200v-20.69L76.69 200Z"/></g>`),
		g.Group(children),
	)
}