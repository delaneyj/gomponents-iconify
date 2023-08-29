package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraSlashDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 64h-32l-16-24H96L80 64H48a16 16 0 0 0-16 16v112a16 16 0 0 0 16 16h160a16 16 0 0 0 16-16V80a16 16 0 0 0-16-16Zm-80 104a36 36 0 1 1 36-36a36 36 0 0 1-36 36Z" opacity=".2"/><path d="M53.92 34.62a8 8 0 1 0-11.84 10.76L51.73 56H48a24 24 0 0 0-24 24v112a24 24 0 0 0 24 24h149.19l4.89 5.38a8 8 0 1 0 11.84-10.76Zm51.66 80.61l37 40.69A27.71 27.71 0 0 1 128 160a28 28 0 0 1-22.42-44.77ZM48 200a8 8 0 0 1-8-8V80a8 8 0 0 1 8-8h18.28l28.41 31.26A44 44 0 0 0 128 176a44.21 44.21 0 0 0 25.44-8.12l29.2 32.12ZM232 80v106a8 8 0 0 1-16 0V80a8 8 0 0 0-8-8h-32a8 8 0 0 1-6.65-3.56L155.71 48h-55.47a8 8 0 0 1-12.91-9.42l2-3A8 8 0 0 1 96 32h64a8 8 0 0 1 6.66 3.56L180.28 56H208a24 24 0 0 1 24 24Z"/></g>`),
		g.Group(children),
	)
}