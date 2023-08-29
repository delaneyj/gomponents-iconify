package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#d33b23" d="M61 39.2c3.7 10.5-4.1 18.4-12.8 19.9c-8.7 1.4-18.4-3.6-18.8-14.7c-.6-14.7 10.1-13 13.9-13.6c6.3-1 13-5 17.7 8.4"/><path fill="none" stroke="#83bf4f" stroke-miterlimit="10" stroke-width="2" d="M45.1 34C43 31.6 34.5 24 37.4 9"/><path fill="#ef4d3c" d="M38.5 42.2C40.7 54.5 30.4 62 20.4 62c-10 0-19.8-7.4-18.2-19.8C4.4 25.7 16 29.6 20.4 29.6c7.2-.1 15.4-3.3 18.1 12.6"/><path fill="none" stroke="#9fc427" stroke-miterlimit="10" stroke-width="2" d="M20.4 33c0-4.1 6-25.5 34.1-30"/><path fill="#ce0f00" d="M13.3 31.6c2.3.3 4.4.7 6.5.9c1.1.1 2.1.2 3.2.3c1.1.1 2.2.1 3.3.1c-1 .6-2.1.9-3.3 1.1c-1.1.2-2.3.2-3.4.1s-2.3-.4-3.3-.8c-1.1-.5-2.2-1-3-1.7"/><path fill="#a51000" d="m38.5 34.8l5.6-1.3c.9-.2 1.8-.5 2.7-.7l2.8-.9c-.6.8-1.4 1.5-2.3 1.9c-.9.5-1.9.9-2.8 1.1c-1 .2-2 .4-3 .4c-1.1 0-2.1-.1-3-.5"/><path fill="#64892f" d="M46.3 7.1c-3.8 8.9 9.3 21.2 9.3 21.2s-1.3-9.3 2.2-17.6C60.1 5.5 54.5 2 54.5 2s-6.3.6-8.2 5.1"/>`),
		g.Group(children),
	)
}