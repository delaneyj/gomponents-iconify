package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSlashDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m120 128l64 48l-64 48Zm0-96v96l64-48Z" opacity=".2"/><path d="m213.92 210.62l-160-176a8 8 0 1 0-11.84 10.76l70.84 77.93L51.2 169.6a8 8 0 1 0 9.6 12.8L112 144v80a8 8 0 0 0 12.8 6.4l50.83-38.12l26.45 29.1a8 8 0 1 0 11.84-10.76ZM128 208v-64l11.73 8.8l25.08 27.59ZM112 71.63V32a8 8 0 0 1 12.8-6.4l64 48a8 8 0 0 1 0 12.8l-33.53 25.15a8 8 0 0 1-9.6-12.8l25-18.75L128 48v23.63a8 8 0 0 1-16 0Z"/></g>`),
		g.Group(children),
	)
}