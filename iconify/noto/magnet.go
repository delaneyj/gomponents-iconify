package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#C62828" d="M67.76 93.24c-17.11 0-30.86-6.73-30.86-22.95s13.65-22.55 30.86-22.55h26.29V16.97H67.76c-30.02 0-55.17 15.08-61.54 40.54H4.65v12.97c0 33.24 28.29 53.51 63.11 53.51h26.29V93.24H67.76z"/><path fill="#82AEC0" d="M94.05 93.24h29.3V124h-29.3zm0-76.27h29.3v30.76h-29.3z"/><path fill="#F44336" d="M67.76 80.26c-17.11 0-30.86-6.73-30.86-22.95s13.65-22.55 30.86-22.55h26.29V4H67.76C32.94 4 4.65 24.28 4.65 57.51c0 33.24 28.29 53.51 63.11 53.51h26.29V80.26H67.76z"/><path fill="#E0E0E0" d="M94.05 80.26h29.3v30.76h-29.3zm0-76.26h29.3v30.76h-29.3z"/>`),
		g.Group(children),
	)
}