package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOfDavid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m12.825 21.5l2.309 4a1 1 0 0 0 1.732 0l2.31-4h4.618a1 1 0 0 0 .866-1.5l-2.31-4l2.31-4a1 1 0 0 0-.866-1.5h-4.619l-2.309-4a1 1 0 0 0-1.732 0l-2.31 4H8.207A1 1 0 0 0 7.34 12l2.31 4l-2.31 4a1 1 0 0 0 .866 1.5h4.619Zm1.732 0h2.886L16 24l-1.443-2.5ZM13.69 20l-2.31-4l2.31-4h4.618l2.31 4l-2.31 4H13.69Zm-3.176-2.5L11.96 20H9.072l1.443-2.5Zm0-3L9.072 12h2.887l-1.444 2.5Zm4.042-4L16 8l1.443 2.5h-2.886ZM20.04 12h2.887l-1.443 2.5L20.04 12Zm1.444 5.5l1.443 2.5H20.04l1.444-2.5Z"/><path d="M1 6a5 5 0 0 1 5-5h20a5 5 0 0 1 5 5v20a5 5 0 0 1-5 5H6a5 5 0 0 1-5-5V6Zm5-3a3 3 0 0 0-3 3v20a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6Z"/></g>`),
		g.Group(children),
	)
}