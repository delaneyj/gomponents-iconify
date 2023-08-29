package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m231.93 97.06l-13.87 104a8 8 0 0 1-7.93 6.94H45.87a8 8 0 0 1-7.93-6.94l-13.87-104A8 8 0 0 1 32 88h192a8 8 0 0 1 7.93 9.06Z" opacity=".2"/><path d="M136 120v56a8 8 0 0 1-16 0v-56a8 8 0 0 1 16 0Zm103.86-21.89L226 202.12A16 16 0 0 1 210.13 216H45.87A16 16 0 0 1 30 202.12l-13.87-104A16 16 0 0 1 32 80h36.37L122 18.73a8 8 0 0 1 12 0L187.63 80H224a16 16 0 0 1 15.85 18.11ZM89.63 80h76.74L128 36.15ZM224 96H32l13.87 104h164.26Zm-51.16 23.2l-5.6 56a8 8 0 0 0 7.16 8.8a7.44 7.44 0 0 0 .81 0a8 8 0 0 0 7.95-7.2l5.6-56a8 8 0 0 0-15.92-1.6Zm-89.68 0a8 8 0 0 0-15.92 1.6l5.6 56a8 8 0 0 0 8 7.2a7.44 7.44 0 0 0 .81 0a8 8 0 0 0 7.16-8.76Z"/></g>`),
		g.Group(children),
	)
}