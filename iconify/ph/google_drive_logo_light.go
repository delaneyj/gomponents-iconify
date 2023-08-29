package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleDriveLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M235.9 152.82L167.43 40.58A13.91 13.91 0 0 0 155.56 34h-55.13a13.92 13.92 0 0 0-11.87 6.58v.05L20.13 152.76a14 14 0 0 0-.28 14.37L47.17 215a14 14 0 0 0 12.15 7h137.35a14 14 0 0 0 12.16-7l27.32-47.82a14 14 0 0 0-.25-14.36ZM222.56 154h-51.17L135 93.33l25.08-41.79Zm-124 0L128 105l29.4 49Zm66 12L191 210H65l26.4-44ZM128 81.67L106.6 46h42.8ZM95.93 51.54L121 93.33L84.6 154H33.43ZM33 166h44.4l-22.73 37.89Zm168.3 37.89L178.59 166H223Z"/>`),
		g.Group(children),
	)
}