package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Accordion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14.75 5.9V3.49L11 2.5L9.44 3L8 2.51L6.57 3L5 2.5l-2.82.75a1.24 1.24 0 0 0-.93 1.2V5.9A1.25 1.25 0 0 0 0 7.15v1.5A1.25 1.25 0 0 0 1.25 9.9v1.65a1.24 1.24 0 0 0 .93 1.2L5 13.5l1.57-.45l1.43.44l1.43-.44l1.57.45l3.75-1V9.9A1.25 1.25 0 0 0 16 8.65v-1.5a1.25 1.25 0 0 0-1.25-1.25zm-11 6l-1.25-.35v-7.1l1.25-.33zm1.25.3V3.8l1 .27v7.86zm2.21-.26V4.06L8 3.82l.79.24v7.88l-.79.24zm2.84 0V4.07L11 3.8v8.4zm3.45-.38l-1.25.33V4.12l1.25.33z"/>`),
		g.Group(children),
	)
}