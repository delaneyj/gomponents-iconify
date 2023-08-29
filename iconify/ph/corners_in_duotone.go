package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornersInDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 96h-48V48ZM96 208v-48H48Zm64 0l48-48h-48ZM96 96V48L48 96Z" opacity=".2"/><path d="M160 104h48a8 8 0 0 0 5.66-13.66l-48-48A8 8 0 0 0 152 48v48a8 8 0 0 0 8 8Zm8-36.69L188.69 88H168ZM96 152H48a8 8 0 0 0-5.66 13.66l48 48A8 8 0 0 0 104 208v-48a8 8 0 0 0-8-8Zm-8 36.69L67.31 168H88ZM208 152h-48a8 8 0 0 0-8 8v48a8 8 0 0 0 13.66 5.66l48-48A8 8 0 0 0 208 152Zm-40 36.69V168h20.69ZM99.06 40.61a8 8 0 0 0-8.72 1.73l-48 48A8 8 0 0 0 48 104h48a8 8 0 0 0 8-8V48a8 8 0 0 0-4.94-7.39ZM88 88H67.31L88 67.31Z"/></g>`),
		g.Group(children),
	)
}