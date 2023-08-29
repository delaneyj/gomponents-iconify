package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.278 1.756c-1.143-.739-2.438-.622-3.417.048c-.967.66-1.611 1.841-1.611 3.229v13.934c0 1.388.644 2.568 1.61 3.23c.98.669 2.275.786 3.418.048l10.789-6.968c1.15-.742 1.683-2.043 1.683-3.277s-.533-2.535-1.683-3.277L6.277 1.755ZM2.75 5.033c0-.921.423-1.625.958-1.991c.522-.358 1.162-.41 1.756-.026l10.789 6.967c.637.41.997 1.18.997 2.017c0 .836-.36 1.606-.997 2.017L5.464 20.985c-.594.383-1.234.33-1.756-.027c-.535-.365-.958-1.07-.958-1.99V5.032Z" clip-rule="evenodd"/><path d="M22.75 5a.75.75 0 0 0-1.5 0v14a.75.75 0 0 0 1.5 0V5Z"/></g>`),
		g.Group(children),
	)
}