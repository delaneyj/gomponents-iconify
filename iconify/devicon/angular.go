package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#DD0031" d="m64 15.36l-47.668 17l7.27 63.027L64 117.762l40.398-22.375l7.27-63.028Zm0 0"/><path fill="#C3002F" d="M64 15.36v11.367v-.051v91.086l40.398-22.375l7.27-63.028Zm0 0"/><path fill="#FFF" d="M64 26.676L34.203 93.492h11.11L51.3 78.54h25.293l5.992 14.953h11.11Zm8.703 42.648H55.297L64 48.383Zm0 0"/>`),
		g.Group(children),
	)
}