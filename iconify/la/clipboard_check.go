package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2c-1.258 0-2.152.89-2.594 2H5v25h22V4h-8.406C18.152 2.89 17.258 2 16 2zm0 2c.555 0 1 .445 1 1v1h3v2h-8V6h3V5c0-.555.445-1 1-1zM7 6h3v4h12V6h3v21H7zm14.281 7.281L15 19.562l-3.281-3.28l-1.438 1.437l4 4l.719.687l.719-.687l7-7z"/>`),
		g.Group(children),
	)
}