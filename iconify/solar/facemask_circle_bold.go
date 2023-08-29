package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacemaskCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m21.824 13.878l-4.112 1.645l-.38 4.938a10.008 10.008 0 0 0 4.492-6.583Zm-6.058 7.389l.444-5.776l-2.632-1.052a4.25 4.25 0 0 0-3.156 0L7.79 15.49l.444 5.776A9.972 9.972 0 0 0 12 22c1.332 0 2.603-.26 3.766-.733Zm-9.098-.806l-.379-4.922l-4.077-1.482a10.011 10.011 0 0 0 4.456 6.404Z"/><path fill-rule="evenodd" d="M21.998 12.193L22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 .13.002.258.007.386l4.98 1.81l2.877-1.15a5.75 5.75 0 0 1 4.271 0L17 14.192l4.998-1.999ZM16 10.5c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5ZM9 12c.552 0 1-.672 1-1.5S9.552 9 9 9s-1 .672-1 1.5s.448 1.5 1 1.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}