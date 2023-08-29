package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bahamut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 15.912a3.456 3.456 0 0 0-3.013-1.08a6.602 6.602 0 0 1-.569 1.478c-.454.91-3.127 2.047-3.127 2.047c-1.136-.796 1.82-7.05 1.82-7.05c-6.282-3.038-20.496-6.329-33.145-3.582a74.344 74.344 0 0 1 11.03 2.331S7.854 10.91 4.5 11.705c7.447.455 21.433.625 13.701 8.528c7.902-2.957 10.006-2.843 8.244 2.501c2.842-3.184 8.684-4.008 7.817-1.961s-5.216 3.965-9.267 3.965c-2.928 0-4.15-2.26-10.446-.924"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.007 20.233c-.151-.645-1.156-.55-2.35-.417c-.34-2.198.944-2.697 2.119-2.526m-2.119 2.526c-2.87.502-4.064 1.639-3.98 2.947s-1.932 2.672-3.496 2.814s-1.023 1.762-1.023 1.762"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.924 22.979c2.082-.984 2.65.75 1.257 2.598m7.249 4.408a7.154 7.154 0 1 0-1.672 9.996m3.537-24.907a3.276 3.276 0 0 0-.49 2.253"/>`),
		g.Group(children),
	)
}