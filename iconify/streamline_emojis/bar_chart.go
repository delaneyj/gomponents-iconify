package streamline_emojis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#fff" d="M5.94 5.93h36.12v36.13H5.94Z"/><path fill="#f0f0f0" d="M40.78 5.93H7.22a1.28 1.28 0 0 0-1.28 1.28v3.21a1.28 1.28 0 0 1 1.28-1.28h33.56a1.28 1.28 0 0 1 1.28 1.28V7.21a1.28 1.28 0 0 0-1.28-1.28Z"/><path fill="#45413c" d="M10.61 45.25a13.4 1.75 0 1 0 26.8 0a13.4 1.75 0 1 0-26.8 0Z" opacity=".15"/><path fill="none" stroke="#e0e0e0" stroke-linecap="round" stroke-linejoin="round" d="M5.94 13.2h36.12v21.61H5.94z"/><path fill="none" stroke="#e0e0e0" stroke-linecap="round" stroke-linejoin="round" d="M5.94 20.09h36.12v7.83H5.94z"/><path fill="none" stroke="#e0e0e0" stroke-linecap="round" stroke-linejoin="round" d="M27.73 5.93v36.13h-7.46V5.93z"/><path fill="none" stroke="#e0e0e0" stroke-linecap="round" stroke-linejoin="round" d="M34.82 5.93v36.13H13.17V5.93z"/><path fill="#ff6242" d="M21.88 26.62h4.24a1.28 1.28 0 0 1 1.28 1.28v14.17h-6.8V27.9a1.28 1.28 0 0 1 1.28-1.28Z"/><path fill="#ff866e" d="M26.12 26.62h-4.24a1.28 1.28 0 0 0-1.28 1.28v2.38A1.28 1.28 0 0 1 21.88 29h4.24a1.28 1.28 0 0 1 1.28 1.28V27.9a1.28 1.28 0 0 0-1.28-1.28Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M21.88 26.62h4.24a1.28 1.28 0 0 1 1.28 1.28v14.17h0h-6.8h0V27.9a1.28 1.28 0 0 1 1.28-1.28Z"/><path fill="#6dd627" d="M11.08 22h4.24a1.28 1.28 0 0 1 1.28 1.28v18.79H9.8V23.24A1.28 1.28 0 0 1 11.08 22Z"/><path fill="#9ceb60" d="M15.32 22h-4.24a1.28 1.28 0 0 0-1.28 1.24v2.38a1.28 1.28 0 0 1 1.28-1.28h4.24a1.28 1.28 0 0 1 1.28 1.28v-2.38A1.28 1.28 0 0 0 15.32 22Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M11.08 22h4.24a1.28 1.28 0 0 1 1.28 1.28v18.79h0h-6.8h0V23.24A1.28 1.28 0 0 1 11.08 22Z"/><path fill="#00b8f0" d="M32.71 13.95h4.24a1.28 1.28 0 0 1 1.28 1.28v26.84h-6.8V15.23a1.28 1.28 0 0 1 1.28-1.28Z"/><path fill="#4acfff" d="M36.94 14h-4.23a1.28 1.28 0 0 0-1.28 1.28v2.32a1.28 1.28 0 0 1 1.28-1.28h4.23a1.28 1.28 0 0 1 1.28 1.28v-2.37A1.28 1.28 0 0 0 36.94 14Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M32.71 13.95h4.24a1.28 1.28 0 0 1 1.28 1.28v26.84h0h-6.8h0V15.23a1.28 1.28 0 0 1 1.28-1.28Z"/><path fill="none" stroke="#45413c" stroke-linecap="round" stroke-linejoin="round" d="M5.94 5.93h36.12v36.13H5.94Z"/>`),
		g.Group(children),
	)
}