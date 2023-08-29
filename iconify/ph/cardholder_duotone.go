package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardholderDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M224 72v48h-56.81a8 8 0 0 0-7.83 6.4a32 32 0 0 1-62.72 0a8 8 0 0 0-7.83-6.4H32V72a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M208 48H48a24 24 0 0 0-24 24v112a24 24 0 0 0 24 24h160a24 24 0 0 0 24-24V72a24 24 0 0 0-24-24ZM40 96h176v16h-48.81a16 16 0 0 0-15.67 12.81a24 24 0 0 1-47 0A16 16 0 0 0 88.81 112H40Zm8-32h160a8 8 0 0 1 8 8v8H40v-8a8 8 0 0 1 8-8Zm160 128H48a8 8 0 0 1-8-8v-56h48.8a40 40 0 0 0 78.39 0H216v56a8 8 0 0 1-8 8Z"/></g>`),
		g.Group(children),
	)
}