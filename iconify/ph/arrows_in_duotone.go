package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsInDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M192 104h-40V64Zm-88 88v-40H64Zm48 0l40-40h-40Zm-48-88V64l-40 40Z" opacity=".2"/><path d="M152 112h40a8 8 0 0 0 5.66-13.66L183.31 84l30.35-30.34a8 8 0 0 0-11.32-11.32L172 72.69l-14.34-14.35A8 8 0 0 0 144 64v40a8 8 0 0 0 8 8Zm8-28.69L172.69 96H160ZM104 144H64a8 8 0 0 0-5.66 13.66L72.69 172l-30.35 30.34a8 8 0 0 0 11.32 11.32L84 183.31l14.34 14.35A8 8 0 0 0 112 192v-40a8 8 0 0 0-8-8Zm-8 28.69L83.31 160H96Zm87.31-.69l14.35-14.34A8 8 0 0 0 192 144h-40a8 8 0 0 0-8 8v40a8 8 0 0 0 13.66 5.66L172 183.31l30.34 30.35a8 8 0 0 0 11.32-11.32Zm-23.31.69V160h12.69ZM107.06 56.61a8 8 0 0 0-8.72 1.73L84 72.69L53.66 42.34a8 8 0 0 0-11.32 11.32L72.69 84L58.34 98.34A8 8 0 0 0 64 112h40a8 8 0 0 0 8-8V64a8 8 0 0 0-4.94-7.39ZM96 96H83.31L96 83.31Z"/></g>`),
		g.Group(children),
	)
}