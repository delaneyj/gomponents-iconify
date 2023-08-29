package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Packagemanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.236 18.833a1.985 1.985 0 1 1 1.985-1.984a1.984 1.984 0 0 1-1.985 1.984Zm11.551 0a1.985 1.985 0 1 1 1.984-1.984a1.985 1.985 0 0 1-1.985 1.984Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 6.615h0a14.147 14.147 0 0 1 14.147 14.147v2.02h0H9.853h0v-2.02A14.147 14.147 0 0 1 24 6.615ZM11.784 4.5l4.019 4.731M36.216 4.5l-4.019 4.731"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.853 22.781h28.294v18.282A2.438 2.438 0 0 1 35.71 43.5H12.29a2.438 2.438 0 0 1-2.437-2.438v-18.28h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.058 33.796a4.447 4.447 0 0 0 0-1.33l1.44-1.1a.33.33 0 0 0 .08-.43l-1.36-2.36a.33.33 0 0 0-.41-.15l-1.7.68a5.68 5.68 0 0 0-1.15-.67l-.26-1.8a.35.35 0 0 0-.34-.29h-2.7a.35.35 0 0 0-.36.27l-.25 1.8a5.483 5.483 0 0 0-1.15.67l-1.7-.68a.33.33 0 0 0-.41.15l-1.37 2.36a.34.34 0 0 0 .09.43l1.43 1.11a4.844 4.844 0 0 0 0 .66v.67l-1.43 1.12a.35.35 0 0 0-.09.44l1.39 2.37a.32.32 0 0 0 .41.14l1.7-.68a4.88 4.88 0 0 0 1.15.67l.25 1.81a.34.34 0 0 0 .34.28h2.66a.34.34 0 0 0 .34-.28l.3-1.81a5.094 5.094 0 0 0 1.15-.67l1.69.68a.35.35 0 0 0 .42-.14l1.36-2.36a.35.35 0 0 0-.08-.44Zm-5.06.84a1.5 1.5 0 1 1 0-3h.01a1.49 1.49 0 0 1 1.49 1.49v.01a1.5 1.5 0 0 1-1.5 1.5Z"/>`),
		g.Group(children),
	)
}