package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StaffSnake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m222.6 43.2l-.1 4.8H288c53 0 96 43 96 96s-43 96-96 96h-40v-80h40c8.8 0 16-7.2 16-16s-7.2-16-16-16h-68l-4.5 144H256c53 0 96 43 96 96s-43 96-96 96h-16v-80h16c8.8 0 16-7.2 16-16s-7.2-16-16-16h-43l-3.1 99.5l-1.4 43.5v1c-.3 8.9-7.6 16-16.5 16s-16.2-7.1-16.5-16v-1l-1-31H136c-22.1 0-40-17.9-40-40s17.9-40 40-40h36l-1-32h-19c-53 0-96-43-96-96c0-47.6 34.6-87.1 80-94.7V256c0 8.8 7.2 16 16 16h16.5L164 128h-41.4c-9 18.9-28.3 32-50.6 32H56c-30.9 0-56-25.1-56-56s25.1-56 56-56h105.5l-.1-4.8L161 32v-1.9c.5-16.6 14.1-30 31-30s30.5 13.4 31 30V32l-.4 11.2zM64 112a16 16 0 1 0 0-32a16 16 0 1 0 0 32z"/>`),
		g.Group(children),
	)
}